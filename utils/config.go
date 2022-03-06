package utils

import "github.com/spf13/viper"

type Config struct {
	AWS struct {
		SNS struct {
			ACCESS_KEY  string `mapstructure:"ACCESS_KEY"`
			SECRET_KEY  string `mapstructure:"SECRET_KEY"`
			REGION      string `mapstructure:"REGION"`
			MAX_RETRIES int    `mapstructure:"MAX_RETRIES"`
			TIMEOUT     int    `mapstructure:"TIMEOUT"`
		}
		SQS struct {
			ACCESS_KEY  string `mapstructure:"ACCESS_KEY"`
			SECRET_KEY  string `mapstructure:"SECRET_KEY"`
			REGION      string `mapstructure:"REGION"`
			MAX_RETRIES int    `mapstructure:"MAX_RETRIES"`
			TIMEOUT     int    `mapstructure:"TIMEOUT"`
		}
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
