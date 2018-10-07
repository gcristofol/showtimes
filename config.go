package main

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for our program
type Config struct {
	DatabaseType     string
	ConnectionString string
	LogFile          string
}

// NewConfig creates a Config instance
func NewConfig() *Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("")
	viper.SetDefault("databasetype", "mssql")
	viper.SetDefault("logfile", "<undefined>")

	viper.BindEnv("databasetype")
	viper.BindEnv("connectionstring")
	viper.BindEnv("logfile")

	databasetype := viper.GetString("databasetype")
	connectionstring := viper.GetString("connectionstring")
	logfile := viper.GetString("logfile")

	cnf := Config{
		DatabaseType:     databasetype,
		ConnectionString: connectionstring,
		LogFile:          logfile,
	}
	return &cnf
}
