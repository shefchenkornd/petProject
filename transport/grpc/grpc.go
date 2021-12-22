package transport

import (
	"context"
	"petProject/pkg/api/grpc"
)

// GRPCServer ...
type GRPCServer struct {
	grpc.UnimplementedUserServiceServer
}

// Add ...
func (G GRPCServer) Add(ctx context.Context, person *grpc.AddPerson) (*grpc.AddResponse, error) {
	var isSuccess bool
	if person.GetName() != "" {
		isSuccess = true
	}

	return &grpc.AddResponse{Success: isSuccess}, nil
}