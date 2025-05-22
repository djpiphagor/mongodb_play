package mongodb

import (
	"context"
	"log/slog"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, connString string) (*mongo.Client, error) {
	ctxWTO, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cl, err := mongo.Connect(ctxWTO, options.Client().ApplyURI(connString))
	if err != nil {
		return nil, errors.Wrap(err, "canot connect to mongo db")
	}
	return cl, nil
}

func MakeConnString(user, pass string, host string, port uint) string {
	conn := url.URL{
		Scheme: "mongodb",
		User:   url.UserPassword(user, pass),
		Host:   net.JoinHostPort(host, strconv.Itoa(int(port))),
	}
	slog.Info(conn.String())
	return conn.String()
}
