package tests

import (
	"encoding/json"
	"github.com/appleboy/gofight/v2"
	"github.com/fufu-yedek/getir-challange/apihelper/response"
	"github.com/fufu-yedek/getir-challange/bunt"
	"github.com/fufu-yedek/getir-challange/memrecords"
	"github.com/fufu-yedek/getir-challange/server"
	"github.com/tidwall/buntdb"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateOrUpdateRecords(t *testing.T) {
	if err := bunt.Initialize(); err != nil {
		t.Fatal("could not initialize in-memory db")
	}

	// prepare
	err := bunt.DB().Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("exist", "value", nil)
		return err
	})

	if err != nil {
		t.Fatalf("error while preparing data %v", err)
	}

	defer func() {
		_ = bunt.DB().Close()
	}()

	mux := server.InitializeRoutesForTest()

	type params struct {
		Key   string
		Value string
	}

	tests := []struct {
		name     string
		params   params
		code     int
		response interface{}
	}{
		{
			name: "it should create new record",
			params: params{
				Key:   "active-tabs",
				Value: "getir",
			},
			code: http.StatusOK,
			response: memrecords.SingleRecordResponse{
				Key:   "active-tabs",
				Value: "getir",
			},
		},
		{
			name: "it should update existing record",
			params: params{
				Key:   "exist",
				Value: "value-updated",
			},
			code: http.StatusOK,
			response: memrecords.SingleRecordResponse{
				Key:   "exist",
				Value: "value-updated",
			},
		},
		{
			name: "it should return error if key field is empty",
			params: params{
				Key:   "",
				Value: "value",
			},
			code: http.StatusBadRequest,
			response: response.BaseResponse{
				Code: 400,
				Msg:  "key field must be filled",
			},
		},
		{
			name: "it should return error if value field is empty",
			params: params{
				Key:   "key",
				Value: "",
			},
			code: http.StatusBadRequest,
			response: response.BaseResponse{
				Code: 400,
				Msg:  "value field must be filled",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp gofight.HTTPResponse

			gofight.New().
				POST("/in-memory").
				SetJSON(gofight.D{
					"key":   tt.params.Key,
					"value": tt.params.Value,
				}).
				Run(mux, func(httpResponse gofight.HTTPResponse, request gofight.HTTPRequest) {
					resp = httpResponse
				})

			if resp.Code != tt.code {
				t.Errorf("expected code %v, but got %v", tt.code, resp.Code)
			}

			expectedBody, err := json.Marshal(tt.response)
			if err != nil {
				t.Errorf("error should be nil")
			}

			gotBody := resp.Body.String()

			if !reflect.DeepEqual(string(expectedBody), gotBody) {
				t.Errorf("expected body %v\nbut got %v", string(expectedBody), gotBody)
			}

			if tt.code == http.StatusOK {
				var gotVal string
				err := bunt.DB().View(func(tx *buntdb.Tx) error {
					val, err := tx.Get(tt.params.Key)
					if err != nil {
						return err
					}

					gotVal = val
					return nil
				})

				if err != nil {
					t.Errorf("error must be nil %v", err)
				}

				if gotVal != tt.params.Value {
					t.Errorf("expected value %v, but got %v", tt.params.Value, gotVal)
				}
			}

		})
	}
}

func TestRetrieve(t *testing.T) {
	if err := bunt.Initialize(); err != nil {
		t.Fatal("could not initialize in-memory db")
	}

	// prepare
	err := bunt.DB().Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("active-tabs", "getir", nil)
		return err
	})

	if err != nil {
		t.Fatalf("error while preparing data %v", err)
	}

	defer func() {
		_ = bunt.DB().Close()
	}()

	mux := server.InitializeRoutesForTest()

	type params struct {
		Key string
	}

	tests := []struct {
		name     string
		params   params
		code     int
		response interface{}
	}{
		{
			name: "it should return existed data",
			params: params{
				Key: "active-tabs",
			},
			code: http.StatusOK,
			response: memrecords.SingleRecordResponse{
				Key:   "active-tabs",
				Value: "getir",
			},
		},
		{
			name: "it should return error if key not found in db",
			params: params{
				Key: "exist",
			},
			code: http.StatusBadRequest,
			response: response.BaseResponse{
				Code: 400,
				Msg:  "could not find record",
			},
		},
		{
			name: "it should return error if key field is empty",
			params: params{
				Key: "",
			},
			code: http.StatusBadRequest,
			response: response.BaseResponse{
				Code: 400,
				Msg:  "key field must be filled",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp gofight.HTTPResponse

			gofight.New().
				GET("/in-memory").
				SetQuery(gofight.H{
					"key": tt.params.Key,
				}).
				Run(mux, func(httpResponse gofight.HTTPResponse, request gofight.HTTPRequest) {
					resp = httpResponse
				})

			if resp.Code != tt.code {
				t.Errorf("expected code %v, but got %v", tt.code, resp.Code)
			}

			expectedBody, err := json.Marshal(tt.response)
			if err != nil {
				t.Errorf("error should be nil")
			}

			gotBody := resp.Body.String()

			if !reflect.DeepEqual(string(expectedBody), gotBody) {
				t.Errorf("expected body %v\nbut got %v", string(expectedBody), gotBody)
			}

		})
	}
}
