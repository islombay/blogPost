package configs

import (
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

type ConfigGrpc struct {
	Server configServer
	DB     configDB
}

type configServer struct {
	Port string
}

type configDB struct {
	Host    string
	Port    string
	DBName  string
	SSLMode string
}

const (
	configFilePathGRPC = "configs"
	configFileNameGRPC = "server"
)

func NewConfigGRPC() ConfigGrpc {
	if err := initConfigGRPCYml(); err != nil {
		slog.Error("could not load grpc yml", sl.Err(err))
		os.Exit(1)
	}
	r := ConfigGrpc{
		Server: configServer{
			Port: viper.GetString("server.port"),
		},
		DB: configDB{
			Host:    viper.GetString("db.host"),
			Port:    viper.GetString("db.port"),
			DBName:  viper.GetString("db.dbname"),
			SSLMode: viper.GetString("db.sslmode"),
		},
	}
	return r
}

func initConfigGRPCYml() error {
	viper.AddConfigPath(configFilePathGRPC)
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFileNameGRPC)
	return viper.ReadInConfig()
}
