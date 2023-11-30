package post_server

import (
	post_service "github.com/islombay/blogPost/internal/grpc/protos"
	"github.com/islombay/blogPost/internal/service/post"
)

type PostServer struct {
	post_service.UnimplementedPostServiceServer
	service post.PostServiceInterface
}

func NewServer(s post.PostServiceInterface) *PostServer {
	return &PostServer{service: s}
}
