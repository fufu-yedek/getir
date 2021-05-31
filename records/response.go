package records

import (
	apiresponse "github.com/fufuceng/getir-challange/apihelper/response"
	"github.com/fufuceng/getir-challange/gtime"
)

type SingleRecordResponse struct {
	// example: TAKwGc6Jr4i8Z487
	Key string `json:"key"`
	// example: 2017-01-28T01:22:14.398Z
	CreatedAt gtime.JSONTime `json:"createdAt"`
	// example: 2800
	TotalCount int `json:"totalCount"`
}

type ListRecordsResponse struct {
	apiresponse.BaseResponse
	Records []SingleRecordResponse `json:"records"`
}

type ListRecordsSerializer struct {
	Records []RecordWithCount
}

func (s ListRecordsSerializer) Response() interface{} {
	resp := ListRecordsResponse{
		BaseResponse: apiresponse.DefaultSuccessResponse,
		Records:      []SingleRecordResponse{},
	}

	for _, record := range s.Records {
		resp.Records = append(resp.Records, SingleRecordResponse{
			Key:        record.Key,
			CreatedAt:  gtime.JSONTime(record.CreatedAt),
			TotalCount: record.TotalCount,
		})
	}

	return resp
}
