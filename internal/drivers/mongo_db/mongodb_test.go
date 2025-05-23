package mongodb_test

import (
	"context"
	"mongodb_play/internal/config"
	mongodb "mongodb_play/internal/drivers/mongo_db"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Test_New(t *testing.T) {
	cfg, err := config.LoadConfig("../../../config/test.yml")
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoDB.ConnTimeout)
	t.Cleanup(
		func() {
			t.Helper()
			cancel()
		})

	connString := mongodb.MakeConnString(
		cfg.MongoDB.Username,
		cfg.MongoDB.Password,
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
	)
	mongoClient, err := mongodb.New(ctx, connString, cfg.MongoDB.ConnTimeout)
	defer mongoClient.Disconnect(ctx)
	require.NoError(t, err)
	err = mongoClient.Ping(ctx, readpref.Primary())
	require.NoError(t, err)
}
