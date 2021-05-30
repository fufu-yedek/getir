package gtime

import (
	"reflect"
	"testing"
	"time"
)

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		t       JSONTime
		args    args
		wantErr bool
		want    JSONTime
	}{
		{
			name: "it should set time if its in expected format",
			t:    JSONTime{},
			args: args{
				b: []byte("2015-01-07"),
			},
			wantErr: false,
			want:    JSONTime(time.Date(2015, 1, 7, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "it should return an error if given time is not in expected format",
			t:    JSONTime{},
			args: args{
				b: []byte("2017-01-28T01:22:14.398Z"),
			},
			wantErr: true,
		},
		{
			name: "it should not update the time if given value is null",
			t:    JSONTime{},
			args: args{
				b: []byte("null"),
			},
			wantErr: false,
			want:    JSONTime{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !tt.t.ToTime().Equal(tt.want.ToTime()) {
					t.Errorf("time values must be equal, got = %v, want = %v", tt.t.ToTime(), tt.want.ToTime())
				}
			}
		})
	}
}

func TestJSONTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		t       JSONTime
		want    []byte
		wantErr bool
	}{
		{
			name:    "it should format time according to the expected format",
			t:       JSONTime(time.Date(2015, 1, 7, 12, 11, 10, 154000000, time.UTC)),
			want:    []byte(`"2015-01-07T12:11:10.154Z"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
