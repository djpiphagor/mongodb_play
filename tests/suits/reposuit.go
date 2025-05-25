package suits

import (
	"context"
	"math/rand/v2"
	"mongodb_play/internal/config"
	"mongodb_play/internal/domain/entities"
	mongodb "mongodb_play/internal/drivers/mongo_db"
	"strconv"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBSuit struct {
	*testing.T
	Config *config.Config
	Client *mongo.Client
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
		T:      t,
		Config: cfg,
		Client: mongoConn,
	}
}

func GeterateRandomCar(t *testing.T) entities.Car {
	t.Helper()

	// generate fake car number
	sb := &strings.Builder{}
	sb.WriteString(
		strings.ToUpper(
			gofakeit.Letter(),
		),
	)
	sb.WriteString(
		strconv.Itoa(
			gofakeit.Number(100, 999),
		),
	)
	sb.WriteString(
		strings.ToUpper(
			gofakeit.LetterN(2),
		),
	)
	sb.WriteString(
		strconv.Itoa(
			gofakeit.Number(10, 777),
		),
	)

	// fake car
	var c entities.Car
	fakeCar := gofakeit.Car()
	c.Brand = fakeCar.Brand
	c.Model = fakeCar.Model
	c.EnginePower = 100 * rand.Float32()
	c.NumberPlate = sb.String()

	// random engine type
	switch {
	case c.EnginePower > 80:
		c.Engine = entities.Gasoline
	case c.EnginePower > 50:
		c.Engine = entities.Diesel
	case c.EnginePower > 20:
		c.Engine = entities.EV
	}

	return c
}
