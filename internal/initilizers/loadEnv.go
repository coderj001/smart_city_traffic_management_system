package initilizers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AutomaticEnv()

	viper.SetEnvPrefix("POSTGRES")
	viper.BindEnv("HOST")
	viper.BindEnv("USER")
	viper.BindEnv("PASSWORD")
	viper.BindEnv("DB")
	viper.BindEnv("PORT")

	config.DBHost = viper.GetString("HOST")
	config.DBUserName = viper.GetString("USER")
	config.DBUserPassword = viper.GetString("PASSWORD")
	config.DBName = viper.GetString("DB")
	config.DBPort = viper.GetString("PORT")

	return config, nil
}
