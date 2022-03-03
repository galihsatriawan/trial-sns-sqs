package utils

import "github.com/spf13/viper"

type Config struct {
	AWS struct {
		ACCESS_KEY string `mapstructure:"AWS_ACCESS_KEY"`
		SECRET_KEY string `mapstructure:"AWS_SECRET_KEY"`
		REGION     string `mapstructure:"AWS_REGION"`
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
