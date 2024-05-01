package config

import (
	"github.com/spf13/viper"
	"log"
)

type Server struct {
	Port int `yaml:"port"`
}

type Config struct {
	Server Server `yaml:"server"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Panicln(err)
	}
	return &cfg
}
