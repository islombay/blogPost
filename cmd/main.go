package main

import (
	"context"
	"github.com/islombay/blogPost"
	"github.com/islombay/blogPost/internal/database/postgres"
	"github.com/islombay/blogPost/internal/service"
	"github.com/islombay/blogPost/pkg/httpserver/handlers"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/islombay/blogPost/configs"
	"github.com/joho/godotenv"
)

func main() {
	cfg := configs.InitConfig()
	slog.Info("yml files are loaded")

	if err := godotenv.Load(); err != nil {
		slog.Error("could not load .env files", sl.Err(err))
	}

	db := postgres.NewPostgresDB(postgres.NewPostgresDBBody{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: os.Getenv("DB_USER"),
		PWD:      os.Getenv("DB_PWD"),
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})

	serv := service.NewBlogPostService(db)

	server := new(blogPost.Server)

	go func() {
		if err := server.Run(cfg.Server.Port, handlers.InitRoutes(serv)); err != nil {
			slog.Error("could not start server", sl.Err(err))
			os.Exit(1)
		}
	}()

	slog.Info("server started")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Info("server stopped")

	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("could not shutdown server", sl.Err(err))
	} else {
		slog.Debug("server shutdown")
	}

	if err := db.Close(); err != nil {
		slog.Error("could not close db", sl.Err(err))
	} else {
		slog.Debug("db closed")
	}
}
