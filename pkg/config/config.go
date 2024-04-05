package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DB_USERNAME  string `mapstructure:"DBUSER"`
	DB_PASSWORD  string `mapstructure:"DBPASS"`
	DB_HOSTNAME  string `mapstructure:"DBHOST"`
	DB_PORT      int    `mapstructure:"DBPORT"`
	DB_NAME      string `mapstructure:"DBNAME"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return c, nil
}
