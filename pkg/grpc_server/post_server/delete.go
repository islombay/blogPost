package post_server

import (
	"context"
	pb "github.com/islombay/blogPost/internal/grpc/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PostServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if err := s.service.Delete(req.GetId()); err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &pb.DeleteResponse{Ok: true}, nil
}
