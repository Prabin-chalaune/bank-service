package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/prabin/bank-service/internal/db/sqlc"
	"github.com/prabin/bank-service/pkg/token"
	"github.com/prabin/bank-service/pkg/util"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	config      util.Config
	store       db.Store
	router      *gin.Engine
	tokenMaker  token.Maker
	redisClient *redis.Client
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	server := &Server{
		config:      config,
		store:       store,
		tokenMaker:  tokenMaker,
		redisClient: redisClient,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	r := gin.Default()

	r.POST("/users", server.createUser)
	r.POST("/users/login", server.loginUser)

	authRoutes := r.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
