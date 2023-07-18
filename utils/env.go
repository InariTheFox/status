package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	Directory string
	Params    *viper.Viper
)

func InitEnv() {
	if Params != nil {
		return
	}

	Params = viper.New()
	Params.AutomaticEnv()

	var err error
	defaultDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get current working directory; are you sure the application has permissions? %v", err)
		defaultDir = "."
	}

	Params.SetDefault("SERVER_IP", "0.0.0.0")
	Params.SetDefault("SERVER_PORT", "8080")
	Params.SetDefault("WORKING_DIR", defaultDir)
	Params.SetDefault("LANGUAGE", "en")
	Params.SetDefault("ADMIN_USER", nil)
	Params.SetDefault("ADMIN_EMAIL", nil)
	Params.SetDefault("ADMIN_PASSWORD", nil)
	Params.SetDefault("DB_DRIVER", "postgres")
	Params.SetDefault("DB_HOST", "localhost")

	dbDriver := Params.GetString("DRIVER")
	dbPort := Params.GetInt("DB_PORT")

	if dbPort == 0 && dbDriver != "sqlite" && dbDriver != "sqlite3" {
		if dbDriver == "postgres" {
			Params.SetDefault("DB_PORT", 5432)
		} else if dbDriver == "mysql" {
			Params.SetDefault("DB_PORT", 3306)
		}
	}

	var conn string
	if dbDriver == "postgres" {
		conn = fmt.Sprintf(
			"postgresql://%s:%s@%s:%d",
			Params.GetString("DB_USER"),
			Params.GetString("DB_PASSWORD"),
			Params.GetString("DB_HOST"),
			Params.GetInt("DB_PORT"))
		Params.Set("DB_CONN", conn)
	}

	Directory := Params.GetString("WORKING_DIR")
	Params.SetConfigName("config")
	Params.SetConfigType("yaml")
	Params.AddConfigPath(Directory)
	Params.ReadInConfig()

	Params.AddConfigPath(Directory)
	Params.SetConfigFile(".env")
	Params.ReadInConfig()

	log.Printf("current working directory: %s", Directory)
}
