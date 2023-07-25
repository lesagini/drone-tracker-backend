package utils

import "github.com/spf13/viper"

type Config struct {
	DbSource string `mapstructure:"DB_SOURCE"`
	DbDriver string `mapstructure:"DB_DRIVER"`
	Address  string `mapstructure:"ADDRESS"`
}

func GetConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
