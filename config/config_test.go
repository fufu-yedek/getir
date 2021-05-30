package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		name           string
		envs           map[string]string
		fileReaderFunc func(s string) ([]byte, error)
		expectedConfig Config
		wantErr        bool
	}{
		{
			name:    "it should return error if env variable does not exist",
			wantErr: true,
		},
		{
			name: "it should return error if an error occurs while reading a file",
			envs: map[string]string{envConfigPath: "test-path"},
			fileReaderFunc: func(s string) ([]byte, error) {
				return nil, fmt.Errorf("random error")
			},
			wantErr: true,
		},
		{
			name: "it should parse config correctly",
			envs: map[string]string{envConfigPath: "test-path"},
			fileReaderFunc: func(s string) ([]byte, error) {
				return json.Marshal(Config{
					Server: Server{
						Host: "host",
						Port: "port",
					},
					Mongo: Mongo{
						Uri:  "uri",
						Name: "name",
					},
				})
			},
			expectedConfig: Config{
				Server: Server{
					Host: "host",
					Port: "port",
				},
				Mongo: Mongo{
					Uri:  "uri",
					Name: "name",
				},
				Env: Env{
					ConfigPath: "test-path",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envs {
				_ = os.Setenv(k, v)
			}

			defer func() {
				for k := range tt.envs {
					_ = os.Unsetenv(k)
				}
			}()

			if tt.fileReaderFunc != nil {
				prevFunc := fileReaderFunc
				fileReaderFunc = tt.fileReaderFunc

				defer func() {
					fileReaderFunc = prevFunc
				}()
			}

			if err := Initialize(); (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(innerConfig, tt.expectedConfig) {
					t.Errorf("initialize() got = %v, want %v", innerConfig, tt.expectedConfig)
				}
			}
		})
	}
}
