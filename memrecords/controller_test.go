package memrecords

import (
	"fmt"
	"github.com/fufu-yedek/getir-challange/gerrors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_controller_CreateOrUpdate(t *testing.T) {
	type mockResponse struct {
		Record Record
		Error  error
	}

	type fields struct {
		repository   *MockRepository
		repoParam    Record
		repoResponse mockResponse
	}

	type args struct {
		params CreateOrUpdateParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "it should return newly created record in expected format",
			fields: fields{
				repository: NewMockRepository(gomock.NewController(t)),
				repoParam: Record{
					Key:   "active-tabs",
					Value: "getir",
				},
				repoResponse: mockResponse{
					Record: Record{
						Key:   "active-tabs",
						Value: "getir",
					},
					Error: nil,
				},
			},
			args: args{
				params: CreateOrUpdateParams{
					Body: struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					}(struct {
						Key   string
						Value string
					}{Key: "active-tabs", Value: "getir"}),
				},
			},
			want: SingleRecordResponse{
				Key:   "active-tabs",
				Value: "getir",
			},
			wantErr: false,
		},
		{
			name: "it should return error if repo returns error",
			fields: fields{
				repository: NewMockRepository(gomock.NewController(t)),
				repoParam: Record{
					Key:   "active-tabs",
					Value: "getir",
				},
				repoResponse: mockResponse{
					Error: fmt.Errorf("some-error"),
				},
			},
			args: args{
				params: CreateOrUpdateParams{
					Body: struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					}(struct {
						Key   string
						Value string
					}{Key: "active-tabs", Value: "getir"}),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := controller{
				repository: tt.fields.repository,
			}

			defer func() {
				tt.fields.repository.Ctrl.Finish()
			}()

			tt.fields.repository.EXPECT().
				CreateOrUpdate(tt.fields.repoParam).
				Return(tt.fields.repoResponse.Record, tt.fields.repoResponse.Error)

			got, err := c.CreateOrUpdate(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got == nil {
					t.Errorf("got should not be nil")
					return
				}

				if !reflect.DeepEqual(got.Response(), tt.want) {
					t.Errorf("CreateOrUpdate() got = %v, want %v", got, tt.want)
				}
			}

		})
	}
}

func Test_controller_Retrieve(t *testing.T) {
	type mockResponse struct {
		Record Record
		Error  error
	}

	type fields struct {
		repository   *MockRepository
		repoParam    Filter
		repoResponse mockResponse
	}

	type args struct {
		params RetrieveParams
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "it should return record in expected format",
			fields: fields{
				repository: NewMockRepository(gomock.NewController(t)),
				repoParam: Filter{
					Key: "active-tabs",
				},
				repoResponse: mockResponse{
					Record: Record{
						Key:   "active-tabs",
						Value: "getir",
					},
					Error: nil,
				},
			},
			args: args{
				params: RetrieveParams{
					Key: "active-tabs",
				},
			},
			want: SingleRecordResponse{
				Key:   "active-tabs",
				Value: "getir",
			},
			wantErr: false,
		},
		{
			name: "it should return error if repo returns error",
			fields: fields{
				repository: NewMockRepository(gomock.NewController(t)),
				repoParam: Filter{
					Key: "active-tabs",
				},
				repoResponse: mockResponse{
					Error: gerrors.ErrRecordNotFound,
				},
			},
			args: args{
				params: RetrieveParams{
					Key: "active-tabs",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := controller{
				repository: tt.fields.repository,
			}

			defer func() {
				tt.fields.repository.Ctrl.Finish()
			}()

			tt.fields.repository.EXPECT().
				FindOne(tt.fields.repoParam).
				Return(tt.fields.repoResponse.Record, tt.fields.repoResponse.Error)

			got, err := c.Retrieve(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got == nil {
					t.Errorf("got should not be nil")
					return
				}

				if !reflect.DeepEqual(got.Response(), tt.want) {
					t.Errorf("CreateOrUpdate() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
