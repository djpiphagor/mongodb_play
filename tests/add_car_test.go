package tests_test

import (
	"log/slog"
	"math/rand/v2"
	"mongodb_play/internal/domain/entities"
	addcar "mongodb_play/internal/domain/usecases/add_car"
	mongodb "mongodb_play/internal/drivers/mongo_db"
	"mongodb_play/internal/infrastructure/repo"
	"mongodb_play/tests/suits"
	"strconv"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestAddCar_HappyPass(t *testing.T) {
	ctx, s := suits.New(t)

	mongoConnStr := mongodb.MakeConnString(
		s.Config.MongoDB.Username,
		s.Config.MongoDB.Password,
		s.Config.MongoDB.Host,
		s.Config.MongoDB.Port,
	)
	mongoClient, err := mongodb.New(ctx, mongoConnStr, s.Config.MongoDB.ConnTimeout)
	if err != nil {
		t.Error("can't connect to mongo db")
	}
	mongoRepo := repo.NewRepo(
		mongoClient,
		s.Config.MongoDB.Database,
		s.Config.MongoDB.Collection,
	)

	add := addcar.New(mongoRepo)

	c := GeterateRandomCar(t)
	_, err = add.Action(ctx, c)
	if err != nil {
		slog.Error(err.Error())
	}

	require.NoError(t, err)

	foundCar, err := mongoRepo.FindCarByNumberPlate(ctx, c.NumberPlate)
	require.NoError(t, err)
	assert.Equal(t, c.NumberPlate, foundCar.NumberPlate)
}

func TestAdd2TheSameCars_Failed(t *testing.T) {
	ctx, s := suits.New(t)

	mongoConnStr := mongodb.MakeConnString(
		s.Config.MongoDB.Username,
		s.Config.MongoDB.Password,
		s.Config.MongoDB.Host,
		s.Config.MongoDB.Port,
	)
	mongoClient, err := mongodb.New(ctx, mongoConnStr, s.Config.MongoDB.ConnTimeout)
	if err != nil {
		t.Error("can't connect to mongo db")
	}
	mongoRepo := repo.NewRepo(
		mongoClient,
		s.Config.MongoDB.Database,
		s.Config.MongoDB.Collection,
	)

	add := addcar.New(mongoRepo)

	c := GeterateRandomCar(t)
	_, err = add.Action(ctx, c)
	if err != nil {
		slog.Error(err.Error())
	}

	require.NoError(t, err)

	foundCar, err := mongoRepo.FindCarByNumberPlate(ctx, c.NumberPlate)
	require.NoError(t, err)
	assert.Equal(t, c.NumberPlate, foundCar.NumberPlate)

	_, err = add.Action(ctx, c)
	require.Error(t, err)
}
