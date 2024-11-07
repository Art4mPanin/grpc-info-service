package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	GRPC     GRPCConfig `yaml:"grpc"`
	DBConfig DBConfig   `yaml:"db"`
}
type GRPCConfig struct {
	Port     int           `yaml:"port"`
	AuthPort int           `yaml:"authport"`
	Timeout  time.Duration `yaml:"timeout"`
}
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Database string `yaml:"database"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Sub("grpc").Unmarshal(&config.GRPC); err != nil {
		return config, err
	}

	if err := viper.Sub("db").Unmarshal(&config.DBConfig); err != nil {
		return config, err
	}

	return config, nil
}
