package main

import (
	"context"

	"github.com/lbrictson/status/pkg"
	"github.com/lbrictson/status/pkg/auth"
	"github.com/lbrictson/status/pkg/configuration"
	"github.com/lbrictson/status/pkg/server"
	"github.com/lbrictson/status/pkg/storage"
)

func main() {
	config := configuration.MustReadConfigurationFromEnvVars()
	serverConfig := server.NewServerConfig{
		Port:  config.WebPort,
		Store: storage.MustNewStore(storage.NewStoreConfig{FileLocation: config.DatabaseFileLocation}),
	}
	// Seed a default user if one doesn't exist
	existingUsers, err := serverConfig.Store.ListOperators(context.Background())
	if err != nil {
		panic(err)
	}
	// This means there are no admins, we need to create one
	if len(existingUsers) == 0 {
		err = serverConfig.Store.SaveOperator(context.Background(), &pkg.Operator{
			Email:          config.DefaultOperator,
			HashedPassword: auth.HashAndSalt(config.DefaultPassword),
			Role:           "Admin",
		})
		if err != nil {
			panic(err)
		}
	}
	server.Run(serverConfig)
}
