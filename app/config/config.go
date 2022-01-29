package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Debug bool `mapstructure:"DEBUG"`

	// //! Server
	// SERVER_PORT    string `mapstructure:"SERVER_PORT"`
	// SERVER_TIMEOUT int    `mapstructure:"SERVER_TIMEOUT"`

	// //! MYSQL
	// MYSQL_DB_HOST string `mapstructure:"MYSQL_DB_HOST"`
	// MYSQL_DB_PORT string `mapstructure:"MYSQL_DB_PORT"`
	// MYSQL_DB_USER string `mapstructure:"MYSQL_DB_USER"`
	// MYSQL_DB_PASS string `mapstructure:"MYSQL_DB_PASS"`
	// MYSQL_DB_NAME string `mapstructure:"MYSQL_DB_NAME"`

	// //! MONGO DB
	// MONGO_DB_HOST string `mapstructure:"MONGO_DB_HOST"`
	// MONGO_DB_PORT string `mapstructure:"MONGO_DB_PORT"`
	// MONGO_DB_USER string `mapstructure:"MONGO_DB_USER"`
	// MONGO_DB_PASS string `mapstructure:"MONGO_DB_PASS"`
	// MONGO_DB_NAME string `mapstructure:"MONGO_DB_NAME"`

	// //! OUATH2 GOOGLE
	// GOOGLE_AUTH_CLIENT string `mapstructure:"GOOGLE_AUTH_CLIENT"`
	// GOOGLE_AUTH_SECRET string `mapstructure:"GOOGLE_AUTH_SECRET"`

	// //! OAUTH2 FACEBOOK
	// FACEBOOK_AUTH_CLIENT string `mapstructure:"FACEBOOK_AUTH_CLIENT"`
	// FACEBOOK_AUTH_SECRET string `mapstructure:"FACEBOOK_AUTH_SECRET"`

	// //! MIDTRANS
	// MIDTRANS_SERVER_KEY  string `mapstructure:"MIDTRANS_SERVER_KEY"`
	// MIDTRANS_CLIENT_KEY  string `mapstructure:"MIDTRANS_CLIENT_KEY"`
	// MIDTRANS_MERCHANT_ID string `mapstructure:"MIDTRANS_MERCHANT_ID"`

	// //! JWT
	// JWT_SECRET  string `mapstructure:"JWT_SECRET"`
	// JWT_EXPIRED int    `mapstructure:"JWT_EXPIRED"`

	// //! REDIS
	// REDIS_ENDPOINT string `mapstructure:"REDIS_ENDPOINT"`
	// REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`

	// //! GOOGLE STORAGE
	// GOOGLE_STORAGE_BUCKET_NAME  string `mapstructure:"GOOGLE_STORAGE_BUCKET_NAME"`
	// GOOGLE_STORAGE_PRIVATE_KEY  string `mapstructure:"GOOGLE_STORAGE_PRIVATE_KEY"`
	// GOOGLE_STORAGE_IAM_EMAIL    string `mapstructure:"GOOGLE_STORAGE_IAM_EMAIL"`
	// GOOGLE_STORAGE_EXPIRED_TIME int    `mapstructure:"GOOGLE_STORAGE_EXPIRED_TIME"`

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

	viper.SetConfigName("config")
	viper.SetConfigType("json")
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
