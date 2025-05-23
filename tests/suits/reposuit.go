package suits

import (
	"context"
	"mongodb_play/internal/config"
	mongodb "mongodb_play/internal/drivers/mongo_db"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBSuit struct {
	*testing.T
	Config     *config.Config
	Collection *mongo.Collection
}

func New(t *testing.T) (context.Context, *MongoDBSuit) {
	t.Helper()
	t.Parallel()

	cfg, err := config.LoadConfig("../config/test.yml")
	if err != nil {
		t.Errorf("incorrect config file")
	}
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoDB.ConnTimeout)

	t.Cleanup(
		func() {
			t.Helper()
			cancel()
		})

	mongoConn, err := mongodb.New(
		ctx,
		mongodb.MakeConnString(
			cfg.MongoDB.Username,
			cfg.MongoDB.Password,
			cfg.MongoDB.Host,
			cfg.MongoDB.Port,
		),
		cfg.MongoDB.ConnTimeout,
	)

	if err != nil {
		t.Errorf("can't connect to mongo db")
	}

	return ctx, &MongoDBSuit{
		T:          t,
		Config:     cfg,
		Collection: mongoConn.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection),
	}
}
