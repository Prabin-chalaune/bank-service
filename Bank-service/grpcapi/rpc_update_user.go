package grpcapi

import (
	"context"

	"github.com/prabin/bank-service/pb"
)

// TODO: add unit test for this too
func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}

//TODO: rpc verify email
