package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

type exampleBody struct {
	Field1 string `json:"field_1"`
	Field2 int    `json:"field_2"`
}

type exampleBodyUnknown struct {
	Field1 string `json:"field_1"`
	Field2 int    `json:"field_2"`
	Field3 bool   `json:"field_3"`
}

func TestParseJSON(t *testing.T) {
	bodyByte, _ := json.Marshal(exampleBody{
		Field1: "field1",
		Field2: 2,
	})

	bodyByte2, _ := json.Marshal(exampleBodyUnknown{
		Field1: "field1",
		Field2: 2,
		Field3: true,
	})

	type args struct {
		req  *http.Request
		dest interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{
			name: "it should parse json body correctly",
			args: args{
				req:  &http.Request{Body: ioutil.NopCloser(strings.NewReader(string(bodyByte)))},
				dest: &exampleBody{},
			},
			wantErr: false,
			want:    exampleBody{Field1: "field1", Field2: 2},
		},
		{
			name: "it should return error if body includes unknown fields",
			args: args{
				req:  &http.Request{Body: ioutil.NopCloser(strings.NewReader(string(bodyByte2)))},
				dest: &exampleBody{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseJSON(tt.args.req, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("ParseJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if reflect.TypeOf(tt.args.dest).Kind() == reflect.Ptr {
					tt.args.dest = reflect.ValueOf(tt.args.dest).Elem().Interface()
				}

				if !reflect.DeepEqual(tt.want, tt.args.dest) {
					t.Errorf("response should be same, got = %v, want = %v", tt.args.dest, tt.want)
				}
			}
		})
	}
}

type testQueryParams struct {
	Query1 string `query:"key"`
	Query2 string `query:"-"`
}

func TestParseQuery(t *testing.T) {
	type args struct {
		req  *http.Request
		dest interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    testQueryParams
	}{
		{
			name: "it should parse query params correctly",
			args: args{
				req: &http.Request{
					URL: &url.URL{
						RawQuery: "key=getir",
					},
				},
				dest: &testQueryParams{},
			},
			wantErr: false,
			want:    testQueryParams{Query1: "getir"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseQuery(tt.args.req, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("ParseQuery() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if reflect.TypeOf(tt.args.dest).Kind() == reflect.Ptr {
					tt.args.dest = reflect.ValueOf(tt.args.dest).Elem().Interface()
				}

				if !reflect.DeepEqual(tt.want, tt.args.dest) {
					t.Errorf("response should be same, got = %v, want = %v", tt.args.dest, tt.want)
				}
			}
		})
	}
}
