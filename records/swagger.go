package records

import "github.com/fufu-yedek/getir-challange/apihelper/response"

// listRecordsSwaggerResponse
// swagger:response listRecordsSwaggerResponse
type listRecordsSwaggerResponse struct {
	// in: body
	Body struct {
		ListRecordsResponse
	}
}

// listRecordsSwaggerErrorResponse
// swagger:response listRecordsSwaggerErrorResponse
type listRecordsSwaggerErrorResponse struct {
	// in: body
	Body struct {
		response.SwaggerErrorResponse
	}
}
