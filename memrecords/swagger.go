package memrecords

import "github.com/fufu-yedek/getir-challange/apihelper/response"

// inMemoryResponse
// swagger:response inMemoryResponse
type inMemoryResponse struct {
	//in: body
	Body struct {
		SingleRecordResponse
	}
}

// inMemoryErrorResponse
// swagger:response inMemoryErrorResponse
type inMemoryErrorResponse struct {
	// in: body
	Body struct {
		response.SwaggerErrorResponse
	}
}
