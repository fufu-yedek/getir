package config

import (
	"os"
	"reflect"
	"testing"
)

func Test_parseEnvironmentVariables(t *testing.T) {
	tests := []struct {
		name    string
		want    Env
		setEnv  map[string]string
		wantErr bool
	}{
		{
			name:    "it should parse env variable correctly",
			want:    Env{ConfigPath: "test-path"},
			setEnv:  map[string]string{envConfigPath: "test-path"},
			wantErr: false,
		},
		{
			name:    "it should return an error env var does not exist",
			want:    Env{},
			setEnv:  nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.setEnv {
				_ = os.Setenv(k, v)
			}

			defer func() {
				for k := range tt.setEnv {
					_ = os.Unsetenv(k)
				}
			}()

			got, err := parseEnvironmentVariables()
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEnvironmentVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEnvironmentVariables() got = %v, want %v", got, tt.want)
			}
		})
	}
}
