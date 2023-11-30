package post_server

import (
	"context"
	pb "github.com/islombay/blogPost/internal/grpc/protos"
	"github.com/islombay/blogPost/internal/service/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *PostServer) Create(ctx context.Context, req *pb.CreatePostBody) (*pb.Post, error) {
	t := time.Unix(req.GetCreatedAt().GetSeconds(), int64(req.GetCreatedAt().GetNanos())).UTC()

	model, err := s.service.Create(post.CreateModel{
		Title:     req.GetTitle(),
		Content:   req.GetContent(),
		CreatedAt: t,
		Username:  req.GetUsername(),
	})

	if err != nil {
		switch err.Error() {
		case post.InvalidModel:
			return nil, status.Errorf(codes.InvalidArgument, "Post body in incorrect")
		default:
			return nil, status.Errorf(codes.Internal, "Internal server error")
		}
	}

	resp := &pb.Post{
		Id:        model.ID,
		Title:     model.Title,
		Content:   model.Content,
		CreatedAt: req.CreatedAt,
		Username:  model.Username,
	}
	return resp, nil
}
