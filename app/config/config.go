package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	SECRET_JWT string = ""
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
	JWT_KEY     string

	// GCP_PROJECT_ID  string
	// GCP_BUCKET_NAME string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.JWT_KEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}
	// if val, found := os.LookupEnv("GCP_PROJECT_ID"); found {
	// 	app.GCP_PROJECT_ID = val
	// 	isRead = false
	// }
	// if val, found := os.LookupEnv("GCP_BUCKET_NAME"); found {
	// 	app.GCP_BUCKET_NAME = val
	// 	isRead = false
	// }

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		app.JWT_KEY = viper.Get("JWT_KEY").(string)
		app.DB_USERNAME = viper.Get("DB_USERNAME").(string)
		app.DB_PASSWORD = viper.Get("DB_PASSWORD").(string)
		app.DB_HOSTNAME = viper.Get("DB_HOSTNAME").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DB_PORT").(string))
		app.DB_NAME = viper.Get("DB_NAME").(string)
		// app.GCP_PROJECT_ID = viper.Get("GCP_PROJECT_ID").(string)
		// app.GCP_BUCKET_NAME = viper.Get("GCP_BUCKET_NAME").(string)

	}

	SECRET_JWT = app.JWT_KEY
	return &app
}
