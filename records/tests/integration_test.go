package tests

import (
	"encoding/json"
	"github.com/appleboy/gofight/v2"
	"github.com/fufu-yedek/getir-challange/apihelper/response"
	"github.com/fufu-yedek/getir-challange/config"
	"github.com/fufu-yedek/getir-challange/gtime"
	"github.com/fufu-yedek/getir-challange/mongo"
	"github.com/fufu-yedek/getir-challange/records"
	"github.com/fufu-yedek/getir-challange/server"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestRecords(t *testing.T) {
	cnf := config.Config{
		Mongo: config.Mongo{
			Uri:  "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true",
			Name: "getir-case-study",
		},
	}

	if err := mongo.Initialize(cnf.Mongo); err != nil {
		t.Fatal("could not initialize mongo connection")
	}

	mux := server.InitializeRoutesForTest()

	type params struct {
		MinCount  int
		MaxCount  int
		StartDate string
		EndDate   string
	}

	tests := []struct {
		name     string
		params   params
		code     int
		response interface{}
	}{
		{
			name: "it should return matched values",
			params: params{
				MinCount:  102,
				MaxCount:  108,
				StartDate: "2015-01-11",
				EndDate:   "2015-02-11",
			},
			code: http.StatusOK,
			response: records.ListRecordsResponse{
				BaseResponse: response.BaseResponse{
					Code: 0,
					Msg:  "Success",
				},
				Records: []records.SingleRecordResponse{
					{
						Key:        "OaxaziGK",
						CreatedAt:  gtime.JSONTime(time.Date(2015, 1, 26, 10, 14, 28, 707000000, time.UTC)),
						TotalCount: 108,
					},
					{
						Key:        "vTurbBka",
						CreatedAt:  gtime.JSONTime(time.Date(2015, 2, 9, 0, 59, 13, 734000000, time.UTC)),
						TotalCount: 103,
					},
				},
			},
		},
		{
			name: "it should return empty list if there is no matched values",
			params: params{
				MinCount:  100,
				MaxCount:  102,
				StartDate: "1900-01-11",
				EndDate:   "1900-02-11",
			},
			code: http.StatusOK,
			response: records.ListRecordsResponse{
				BaseResponse: response.BaseResponse{
					Code: 0,
					Msg:  "Success",
				},
				Records: []records.SingleRecordResponse{},
			},
		},
		{
			name: "it should return error if min count grater than max count",
			params: params{
				MinCount: 102,
				MaxCount: 100,
			},
			code: http.StatusBadRequest,
			response: response.BaseResponse{
				Code: 400,
				Msg:  "min_count must be less than or equal to the max_count",
			},
		},
		{
			name: "it should return error if startDate bigger than endDate",
			params: params{
				StartDate: "2015-03-11",
				EndDate:   "2015-02-11",
			},
			code: http.StatusBadRequest,
			response: response.BaseResponse{
				Code: 400,
				Msg:  "start_date must be less than or equal to the end_date",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp gofight.HTTPResponse

			gofight.New().
				POST("/records").
				SetJSON(gofight.D{
					"min_count":  tt.params.MinCount,
					"max_count":  tt.params.MaxCount,
					"start_date": tt.params.StartDate,
					"end_date":   tt.params.EndDate,
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
