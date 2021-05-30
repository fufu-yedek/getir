package memrecords

import "github.com/fufuceng/getir-challange/apihelper/errors"

type CreateOrUpdateParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p CreateOrUpdateParams) Validate() error {
	if p.Key == "" {
		return errors.NewUserReadableErrf("key field must be filled")
	}

	if p.Value == "" {
		return errors.NewUserReadableErrf("value field must be filled")
	}

	return nil
}

type RetrieveParams struct {
	Key string `json:"key" query:"key"`
}

func (p RetrieveParams) Validate() error {
	if p.Key == "" {
		return errors.NewUserReadableErrf("key field must be filled")
	}

	return nil
}
