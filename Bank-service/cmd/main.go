package main

import (
	"context"
	"net"
	"net/http"
	"os"

	grpcapi "github.com/prabin/bank-service/grpcapi"
	// "github.com/prabin/bank-service/pkg/s3"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prabin/bank-service/internal/api"
	db "github.com/prabin/bank-service/internal/db/sqlc"
	logger "github.com/prabin/bank-service/internal/logger"
	_rabbitmq "github.com/prabin/bank-service/internal/rabbitm"
	"github.com/prabin/bank-service/pb"
	"github.com/prabin/bank-service/pkg/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg(`Can't establish connection to the Postgres!`)
	}

	store := db.NewStore(connPool)

	r, err := _rabbitmq.NewRabbitMQClient(config.RabbitMQURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create rabbitMQ client!")
	}
	defer r.Close()

	go runGatewayServer(config, store)
	runGrpcServer(config, store)
	// s3.InitS3(config.AWSBucketName, config.AWSRegion)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start the gRPC server!")
	}

	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterGoBankServer(grpcServer, server)
	reflection.Register(grpcServer) // provides information about publicly-accessible gRPC services on a gRPC server

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}
	log.Info().Msgf("gRPC server started at port: %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server!")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't create server!")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't start the server!")
	}
}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start the server!")
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterGoBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handle server!")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// statikFs, err := fs.New()
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot create statik fs")
	// }
	// swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFs))
	// mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}
	log.Info().Msgf("HTTP gateway server started at port: %s", listener.Addr().String())
	handler := logger.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server!")
	}
}
