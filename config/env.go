package config

import (
	"fmt"
	"os"
)

const envConfigPath = "GETIR_CONFIG_PATH"
const envPort = "PORT"

type Env struct {
	ConfigPath string `json:"config_path"`
	Port       string `json:"port"`
}

func (e Env) Validate() error {
	if e.ConfigPath == "" {
		return fmt.Errorf("%w: config path required", ErrConfigValidation)
	}

	return nil
}

func parseEnvironmentVariables() (Env, error) {
	env := Env{
		ConfigPath: os.Getenv(envConfigPath),
		Port:       os.Getenv(envPort),
	}

	if err := env.Validate(); err != nil {
		return Env{}, err
	}

	return env, nil
}
