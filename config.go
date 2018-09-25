package main

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for our program
type Config struct {
	DatabaseType     string
	ConnectionString string
}

// NewConfig creates a Config instance
func NewConfig() *Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("")
	viper.SetDefault("databasetype", "mssql")

	viper.BindEnv("databasetype")
	viper.BindEnv("connectionstring")

	databasetype := viper.GetString("databasetype")
	connectionstring := viper.GetString("connectionstring")

	cnf := Config{
		DatabaseType:     databasetype,
		ConnectionString: connectionstring,
	}
	return &cnf
}
