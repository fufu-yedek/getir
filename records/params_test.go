package records

import (
	"github.com/fufu-yedek/getir-challange/gtime"
	"testing"
	"time"
)

func TestListRecordParams_Validate(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "it should return an error if min count bigger than max count",
			fields: fields{
				Body: struct {
					StartDate gtime.JSONTime `json:"start_date"`
					EndDate   gtime.JSONTime `json:"end_date"`
					MinCount  uint           `json:"min_count"`
					MaxCount  uint           `json:"max_count"`
				}{
					MinCount: 2,
					MaxCount: 1,
				},
			},
			wantErr: true,
		},
		{
			name: "it should return an error if startDate bigger than beginDate",
			fields: fields{
				Body: struct {
					StartDate gtime.JSONTime `json:"start_date"`
					EndDate   gtime.JSONTime `json:"end_date"`
					MinCount  uint           `json:"min_count"`
					MaxCount  uint           `json:"max_count"`
				}{
					StartDate: gtime.JSONTime(time.Now()),
					EndDate:   gtime.JSONTime(time.Now().Add(-1 * time.Minute)),
				},
			},

			wantErr: true,
		},
		{
			name: "it should not return an error if all params as expected",
			fields: fields{
				Body: struct {
					StartDate gtime.JSONTime `json:"start_date"`
					EndDate   gtime.JSONTime `json:"end_date"`
					MinCount  uint           `json:"min_count"`
					MaxCount  uint           `json:"max_count"`
				}{
					StartDate: gtime.JSONTime(time.Now()),
					EndDate:   gtime.JSONTime(time.Now().Add(1 * time.Minute)),
					MinCount:  1,
					MaxCount:  2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ListRecordParams{
				Body: tt.fields.Body,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
