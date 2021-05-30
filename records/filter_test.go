package records

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
	"time"
)

func TestFilter_Validate(t *testing.T) {
	type fields struct {
		StartDate time.Time
		EndDate   time.Time
		MinCount  uint
		MaxCount  uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "it should return an error if min count bigger than max count",
			fields: fields{
				MinCount: 9,
				MaxCount: 2,
			},
			wantErr: true,
		},
		{
			name: "it should return an error if startDate bigger than beginDate",
			fields: fields{
				StartDate: time.Now(),
				EndDate:   time.Now().Add(-1 * time.Minute),
			},
			wantErr: true,
		},
		{
			name: "it should not return an error if all params as expected",
			fields: fields{
				StartDate: time.Now(),
				EndDate:   time.Now().Add(1 * time.Minute),
				MinCount:  1,
				MaxCount:  2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Filter{
				StartDate: tt.fields.StartDate,
				EndDate:   tt.fields.EndDate,
				MinCount:  tt.fields.MinCount,
				MaxCount:  tt.fields.MaxCount,
			}
			if err := f.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateMongoQuery(t *testing.T) {
	type args struct {
		from Filter
	}
	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "it should return expected query",
			args: args{
				from: Filter{
					StartDate: time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2020, 12, 21, 0, 0, 0, 0, time.UTC),
					MinCount:  1,
					MaxCount:  2,
				},
			},
			want: bson.M{
				"$and": []bson.M{
					{"createdAt": bson.M{"$gte": time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)}},
					{"createdAt": bson.M{"$lte": time.Date(2020, 12, 21, 0, 0, 0, 0, time.UTC)}},
					{"totalCount": bson.M{"$gte": 1}},
					{"totalCount": bson.M{"$lte": 2}},
				},
			},
		},
		{
			name: "it should return nil query if filter is empty",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateMongoQuery(tt.args.from); !reflect.DeepEqual(fmt.Sprintf("%v", got), fmt.Sprintf("%v", tt.want)) {
				t.Errorf("GenerateMongoQuery() = \n\tgot  %v\n\twant %v", got, tt.want)
			}
		})
	}
}
