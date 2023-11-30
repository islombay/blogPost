package post_server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/islombay/blogPost/internal/grpc/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *PostServer) GetAll(ctx context.Context, empty *empty.Empty) (*pb.GetAllResponse, error) {
	models, err := s.service.GetAll()
	if err != nil {
		switch err.Error() {
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}
	resp := pb.GetAllResponse{}
	for _, m := range models {
		resp.Posts = append(resp.Posts, &pb.Post{
			Id:        m.ID,
			Title:     m.Title,
			Content:   m.Content,
			CreatedAt: timestamppb.New(m.CreatedAt),
			Username:  m.Username,
		})
	}
	return &resp, nil
}
