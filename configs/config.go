package configs

import (
	"log"

	"github.com/spf13/viper"
)

const (
	configFilePath = "configs"
	configFileName = "config"
)

type serverYml struct {
	Host string
	Port string
}

type dbYml struct {
	Host    string
	Port    string
	DBName  string
	SSLMode string
}

type Config struct {
	Server serverYml
	DB     dbYml
}

func InitConfig() Config {
	if err := initConfigYml(); err != nil {
		log.Fatalf("could not load yml file: %v", err)
	}

	return Config{
		Server: serverYml{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		DB: dbYml{
			Host:    viper.GetString("db.host"),
			Port:    viper.GetString("db.port"),
			DBName:  viper.GetString("db.dbname"),
			SSLMode: viper.GetString("db.sslmode"),
		},
	}
}

func initConfigYml() error {
	viper.AddConfigPath(configFilePath)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
