package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, connString string) (*mongo.Client, error) {
	ctxWTO, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	if cl, err := mongo.Connect(ctxWTO, options.Client().ApplyURI(connString)); err != nil {
		return nil, err
	} else {
		return cl, nil
	}
}

func MakeConnString(user, pass string, host string, port uint) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pass, host, port)
}
