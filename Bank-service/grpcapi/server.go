package grpcapi

import (
	"fmt"

	db "github.com/prabin/bank-service/internal/db/sqlc"
	"github.com/prabin/bank-service/pb"
	"github.com/prabin/bank-service/pkg/token"
	"github.com/prabin/bank-service/pkg/util"
)

type Server struct {
	pb.UnimplementedGoBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	return server, nil
}
