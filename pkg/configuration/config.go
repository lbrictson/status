package configuration

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// Server configs
	WebPort int `default:"6672"`
	// Auth configs

	// Database configs
}

func MustReadConfigurationFromEnvVars() *Config {
	c := Config{}
	// expects env variable like 'SETTINGS_WEBPORT' --> WebPort
	err := envconfig.Process("SETTINGS", &c)
	if err != nil {
		fmt.Printf("unable to read config from env vars, panicing %v", err.Error())
		panic(err.Error())
	}
	return &c
}
