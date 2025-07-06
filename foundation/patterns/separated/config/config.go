package config

import (
	"github.com/spf13/viper"
)

type ConnectRPC struct {
	Address string
}

type Config struct {
	ConnectRPC ConnectRPC
}

func New() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("CONNECTRPC_ADDRESS", ":8080")

	return &Config{
		ConnectRPC: ConnectRPC{
			Address: viper.GetString("CONNECTRPC_ADDRESS"),
		},
	}
}

func (c *Config) Load() *Config {
	// Config is already loaded in New()
	return c
}
