package records

import (
	"github.com/fufuceng/getir-challange/apihelper/response"
	"github.com/fufuceng/getir-challange/gtime"
	"reflect"
	"testing"
	"time"
)

func TestListRecordsSerializer_Response(t *testing.T) {
	type fields struct {
		Records []RecordWithCount
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "it should return expected response when records array include elements",
			fields: fields{
				Records: []RecordWithCount{
					{
						Key:        "key-1",
						CreatedAt:  time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
						TotalCount: 1,
					},
					{
						Key:        "key-2",
						CreatedAt:  time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
						TotalCount: 2,
					},
				},
			},
			want: ListRecordsResponse{
				BaseResponse: response.BaseResponse{
					Code: 0,
					Msg:  "Success",
				},
				Records: []SingleRecordResponse{
					{
						Key:        "key-1",
						CreatedAt:  gtime.JSONTime(time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)),
						TotalCount: 1,
					},
					{
						Key:        "key-2",
						CreatedAt:  gtime.JSONTime(time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)),
						TotalCount: 2,
					},
				},
			},
		},
		{
			name:   "it should return expected response when records array empty",
			fields: fields{},
			want: ListRecordsResponse{
				BaseResponse: response.BaseResponse{
					Code: 0,
					Msg:  "Success",
				},
				Records: []SingleRecordResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ListRecordsSerializer{
				Records: tt.fields.Records,
			}
			if got := s.Response(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response() = %v, want %v", got, tt.want)
			}
		})
	}
}
