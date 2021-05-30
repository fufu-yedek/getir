package memrecords

import "testing"

func TestCreateOrUpdateParams_Validate(t *testing.T) {
	type fields struct {
		Key   string
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "it should return error if key field is empty",
			fields: fields{
				Key:   "",
				Value: "value-1",
			},
			wantErr: true,
		},
		{
			name: "it should return error if value field is empty",
			fields: fields{
				Key:   "key-1",
				Value: "",
			},
			wantErr: true,
		},
		{
			name: "it should not return error if key and value fields exist",
			fields: fields{
				Key:   "key-1",
				Value: "value-1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateOrUpdateParams{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRetrieveParams_Validate(t *testing.T) {
	type fields struct {
		Key string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "it should return error if key field is empty",
			fields:  fields{Key: ""},
			wantErr: true,
		},
		{
			name:    "it should not return error if key field exist",
			fields:  fields{Key: "key-1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RetrieveParams{
				Key: tt.fields.Key,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
