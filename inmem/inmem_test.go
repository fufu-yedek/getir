package inmem

import "testing"

func TestInitialize(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "it should open buntdb connection correctly",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Initialize(); (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if conn == nil {
					t.Errorf("conn object should not be nil")
				}

				defer func() { _ = conn.Close() }()
			}

		})
	}
}
