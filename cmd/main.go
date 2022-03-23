package main

import (
	"github.com/lbrictson/status/pkg/configuration"
	"github.com/lbrictson/status/pkg/server"
)

func main() {
	config := configuration.MustReadConfigurationFromEnvVars()
	server.Run(config)
}
