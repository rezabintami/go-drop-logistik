package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	onceConfig sync.Once
)

func GetConfiguration(code string) string {
	onceConfig.Do(func() {
		viper.SetConfigName("config.dev")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(os.Getenv("APP_PATH") + "app/config/")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("err GetConfiguration: %s, code: %s", err.Error(), code)
		}
	})

	value := viper.GetString(code)
	if value == "" && !viper.InConfig(code) {
		log.Fatalf("err GetConfiguration: not found, code: %s \n", code)
	}

	return value
}
