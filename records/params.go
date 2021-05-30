package records

import (
	"github.com/fufuceng/getir-challange/apihelper/errors"
	"github.com/fufuceng/getir-challange/gtime"
	"time"
)

type ListRecordParams struct {
	StartDate gtime.JSONTime `json:"start_date"`
	EndDate   gtime.JSONTime `json:"end_date"`
	MinCount  uint           `json:"min_count"`
	MaxCount  uint           `json:"max_count"`

	ParsedStartDate time.Time `json:"-"`
	ParsedEndDate   time.Time `json:"-"`
}

func (p ListRecordParams) Validate() error {
	if p.MinCount > 0 && p.MaxCount > 0 {
		if p.MinCount > p.MaxCount {
			return errors.NewUserReadableErrf("min_count must be less than or equal to the max_count")
		}
	}

	if !p.StartDate.ToTime().IsZero() && !p.EndDate.ToTime().IsZero() {
		if p.StartDate.ToTime().After(p.EndDate.ToTime()) {
			return errors.NewUserReadableErrf("start_date must be less than or equal to the end_date")
		}
	}

	return nil
}
