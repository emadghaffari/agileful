package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	// Global config
	Confs Cnfs = &Config{}
)

type Cnfs interface {
	Get() Config
	SetDebug(debug bool)
	Set(data []byte) error
}

// Config is base of configs we need for project
type Config struct {
	Debug    bool     // if true we run on debug mode
	POSTGRES Database `yaml:"database"`
}

func (g Config) Get() Config {
	return g
}

func (g *Config) SetDebug(debug bool) {
	g.Debug = debug
}

func (g Config) Set(data []byte) error {
	err := yaml.Unmarshal(data, &g)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the config: %s", err.Error())
	}

	return nil
}
