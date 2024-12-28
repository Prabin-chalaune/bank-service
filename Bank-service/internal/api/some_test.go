package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/prabin/bank-service/internal/db/sqlc"
	"github.com/prabin/bank-service/pkg/util"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: 15,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
