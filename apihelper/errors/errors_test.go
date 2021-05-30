package errors

import "testing"

func TestNewUserReadableErrf(t *testing.T) {
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantMsg string
	}{
		{
			name: "it should create an user readable error",
			args: args{
				s:    "args %v %v",
				args: []interface{}{1, 2},
			},
			wantMsg: "args 1 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewUserReadableErrf(tt.args.s, tt.args.args...)

			if err == nil {
				t.Errorf("err sohuld not be nil")
				return
			}

			if err.Error() != tt.wantMsg {
				t.Errorf("messages must be equal, got = %v, want = %v", err.Error(), tt.wantMsg)
			}
		})
	}
}
