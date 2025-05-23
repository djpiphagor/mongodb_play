package mongodb

import (
	"context"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, connString string, timeout time.Duration) (*mongo.Client, error) {
	ctxWTO, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cl, err := mongo.Connect(ctxWTO, options.Client().ApplyURI(connString))
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to mongo db")
	}
	return cl, nil
}

func MakeConnString(user, pass string, host string, port int) string {
	conn := url.URL{
		Scheme: "mongodb",
		User:   url.UserPassword(user, pass),
		Host:   net.JoinHostPort(host, strconv.Itoa(port)),
	}
	return conn.String()
}
