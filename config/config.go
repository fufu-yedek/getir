package config

import (
	"encoding/json"
	"io/ioutil"
)

var innerConfig Config

type Config struct {
	Server Server `json:"server"`
	Mongo  Mongo  `json:"mongo"`
	Env    Env    `json:"-"`
}

type Mongo struct {
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// Initialize parse and build new config instance using environment variables
func Initialize() error {
	env, err := parseEnvironmentVariables()
	if err != nil {
		return err
	}

	byteConfig, err := fileReaderFunc(env.ConfigPath)
	if err != nil {
		return err
	}

	config := Config{Env: env}
	if err := json.Unmarshal(byteConfig, &config); err != nil {
		return err
	}

	if config.Env.Port != "" {
		config.Server.Port = config.Env.Port
	}

	innerConfig = config
	return nil
}

func Get() Config {
	return innerConfig
}

var fileReaderFunc = ioutil.ReadFile
