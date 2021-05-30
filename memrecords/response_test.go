package memrecords

import (
	"reflect"
	"testing"
)

func TestRecordSerializer_Response(t *testing.T) {
	type fields struct {
		Record Record
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "it should return expected response",
			fields: fields{
				Record: Record{
					Key:   "active-tabs",
					Value: "getir",
				},
			},
			want: SingleRecordResponse{
				Key:   "active-tabs",
				Value: "getir",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RecordSerializer{
				Record: tt.fields.Record,
			}

			if got := r.Response(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response() = %v, want %v", got, tt.want)
			}
		})
	}
}
