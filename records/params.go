package records

import (
	"github.com/fufuceng/getir-challange/apihelper/errors"
	"github.com/fufuceng/getir-challange/gtime"
)

//swagger:parameters Records List-Records
type ListRecordParams struct {
	//in: body
	Body struct {
		//example: 2020-12-20
		StartDate gtime.JSONTime `json:"start_date"`
		//example: 2020-11-20
		EndDate gtime.JSONTime `json:"end_date"`
		//example: 30
		MinCount uint `json:"min_count"`
		//example: 50
		MaxCount uint `json:"max_count"`
	}
}

func (p ListRecordParams) Validate() error {
	if p.Body.MinCount > 0 && p.Body.MaxCount > 0 {
		if p.Body.MinCount > p.Body.MaxCount {
			return errors.NewErrUserReadable("min_count must be less than or equal to the max_count")
		}
	}

	if !p.Body.StartDate.ToTime().IsZero() && !p.Body.EndDate.ToTime().IsZero() {
		if p.Body.StartDate.ToTime().After(p.Body.EndDate.ToTime()) {
			return errors.NewErrUserReadable("start_date must be less than or equal to the end_date")
		}
	}

	return nil
}
