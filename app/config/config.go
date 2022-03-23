package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	
	//! APP
	App struct {
		Env string `mapstructure:"env"`
		Debug bool `mapstructure:"debug"`
		Version string `mapstructure:"version"`
	} `mapstructure:"app"`

	//! Server
	Server struct {
		Address string `mapstructure:"address"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"server"`

	//! MYSQL
	Mysql struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	} `mapstructure:"mysql"`

	//! MONGO DB
	Mongo struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	} `mapstructure:"mongo"`

	//! JWT
	JWT struct {
		Secret  string `mapstructure:"secret"`
		Expired int    `mapstructure:"expired"`
	} `mapstructure:"jwt"`
}

func GetConfig() Config {
	var conf Config

	viper.SetConfigName("config.prod")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("APP_PATH") + "app/config/")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error: ", err)
		conf.Mysql.Host = os.Getenv("MYSQL_DB_HOST")
		conf.Mysql.Port = os.Getenv("MYSQL_DB_PORT")
		conf.Mysql.User = os.Getenv("MYSQL_DB_USER")
		conf.Mysql.Pass = os.Getenv("MYSQL_DB_PASS")
		conf.Mysql.Name = os.Getenv("MYSQL_DB_NAME")

		conf.JWT.Secret = os.Getenv("JWT_SECRET")
		conf.JWT.Expired, _ = strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
	return conf
}
