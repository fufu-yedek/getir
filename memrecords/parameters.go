package memrecords

import "github.com/fufu-yedek/getir-challange/apihelper/errors"

//swagger:parameters InMemory CreateOrUpdate
type CreateOrUpdateParams struct {
	// in: body
	Body struct {
		// example: active-tabs
		// required: true
		Key string `json:"key"`
		// example: getir
		// required: true
		Value string `json:"value"`
	}
}

func (p CreateOrUpdateParams) Validate() error {
	if p.Body.Key == "" {
		return errors.NewErrUserReadable("key field must be filled")
	}

	if p.Body.Value == "" {
		return errors.NewErrUserReadable("value field must be filled")
	}

	return nil
}

//swagger:parameters InMemory Retrieve
type RetrieveParams struct {
	// in: query
	// required: true
	// example: active-tabs
	Key string `json:"key" query:"key"`
}

func (p RetrieveParams) Validate() error {
	if p.Key == "" {
		return errors.NewErrUserReadable("key field must be filled")
	}

	return nil
}
