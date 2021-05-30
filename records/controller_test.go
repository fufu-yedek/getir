package records

import (
	"fmt"
	response2 "github.com/fufuceng/getir-challange/apihelper/response"
	"github.com/fufuceng/getir-challange/gerrors"
	"github.com/fufuceng/getir-challange/gtime"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func Test_controller_ListRecords(t *testing.T) {
	type mockResponse struct {
		Records []RecordWithCount
		Error   error
	}

	type fields struct {
		RecordRepository *MockRepository
		Filter           Filter
		Response         mockResponse
	}

	type args struct {
		params ListRecordParams
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        interface{}
		wantErr     bool
		wantErrType interface{}
	}{
		{
			name: "it should return records list in expected format",
			fields: fields{
				RecordRepository: NewMockRepository(gomock.NewController(t)),
				Filter: Filter{
					StartDate: time.Now(),
					EndDate:   time.Now().Add(1 * time.Minute),
					MinCount:  1,
					MaxCount:  5,
				},
				Response: mockResponse{
					Records: []RecordWithCount{
						{
							Key:        "key1",
							CreatedAt:  time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
							TotalCount: 2,
						},
						{
							Key:        "key2",
							CreatedAt:  time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
							TotalCount: 3,
						},
					},
					Error: nil,
				},
			},
			want: ListRecordsResponse{
				BaseResponse: response2.DefaultSuccessResponse,
				Records: []SingleRecordResponse{
					{
						Key:        "key1",
						CreatedAt:  gtime.JSONTime(time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)),
						TotalCount: 2,
					},
					{
						Key:        "key2",
						CreatedAt:  gtime.JSONTime(time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)),
						TotalCount: 3,
					},
				},
			},
		},
		{
			name: "it should return empty list if repo returns ErrRecordNotFound",
			fields: fields{
				RecordRepository: NewMockRepository(gomock.NewController(t)),
				Filter: Filter{
					StartDate: time.Now(),
					EndDate:   time.Now().Add(1 * time.Minute),
					MinCount:  1,
					MaxCount:  5,
				},
				Response: mockResponse{
					Records: nil,
					Error:   gerrors.ErrRecordNotFound,
				},
			},
			want: ListRecordsResponse{
				BaseResponse: response2.DefaultSuccessResponse,
				Records:      []SingleRecordResponse{},
			},
		},
		{
			name: "it should return error if repo returns some internal error",
			fields: fields{
				RecordRepository: NewMockRepository(gomock.NewController(t)),
				Filter: Filter{
					StartDate: time.Now(),
					EndDate:   time.Now().Add(1 * time.Minute),
					MinCount:  1,
					MaxCount:  5,
				},
				Response: mockResponse{
					Records: nil,
					Error:   fmt.Errorf("some-error"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := controller{
				RecordRepository: tt.fields.RecordRepository,
			}

			defer func() {
				tt.fields.RecordRepository.Ctrl.Finish()
			}()

			tt.fields.RecordRepository.EXPECT().FindWithCount(tt.fields.Filter).Return(tt.fields.Response.Records, tt.fields.Response.Error)

			got, err := c.ListRecords(ListRecordParams{
				Body: struct {
					StartDate gtime.JSONTime `json:"start_date"`
					EndDate   gtime.JSONTime `json:"end_date"`
					MinCount  uint           `json:"min_count"`
					MaxCount  uint           `json:"max_count"`
				}(struct {
					StartDate gtime.JSONTime
					EndDate   gtime.JSONTime
					MinCount  uint
					MaxCount  uint
				}{StartDate: gtime.JSONTime(tt.fields.Filter.StartDate), EndDate: gtime.JSONTime(tt.fields.Filter.EndDate), MinCount: tt.fields.Filter.MinCount, MaxCount: tt.fields.Filter.MaxCount}),
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("ListRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got == nil {
					t.Errorf("got should not be nil")
					return
				}

				if !reflect.DeepEqual(got.Response(), tt.want) {
					t.Errorf("ListRecords() \n\tgot  =%v\n\twant =%v", got, tt.want)
				}
			}

		})
	}
}
