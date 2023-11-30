package main

import (
	"fmt"
	"github.com/islombay/blogPost/configs"
	"github.com/islombay/blogPost/internal/database/postgres"
	"github.com/islombay/blogPost/internal/service"
	"github.com/islombay/blogPost/pkg/grpc_server/post_server"
	"github.com/islombay/blogPost/pkg/utils/logger/sl"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"

	pb "github.com/islombay/blogPost/internal/grpc/protos"
)

func main() {
	cfg := configs.NewConfigGRPC()

	lis, err := net.Listen("tcp", fmt.Sprintf("[::1]:%s", cfg.Server.Port))
	if err != nil {
		slog.Error("could not listen", sl.Err(err))
	}
	server := grpc.NewServer()

	if err := godotenv.Load(); err != nil {
		slog.Error("could not get .env variables", sl.Err(err))
	}

	db := postgres.NewPostgresDB(postgres.NewPostgresDBBody{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: os.Getenv("DB_USER"),
		PWD:      os.Getenv("DB_PWD"),
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})

	services := service.NewBlogPostService(db)

	postService := post_server.NewServer(services.Post)
	pb.RegisterPostServiceServer(server, postService)

	slog.Info(fmt.Sprintf("Running GRPC server on %s", lis.Addr()))
	if err := server.Serve(lis); err != nil {
		slog.Error("could not serve", sl.Err(err))
		os.Exit(1)
	}
}
