package tests

import (
	"log/slog"
	"math/rand/v2"
	"mongodb_play/internal/domain/entities"
	addcar "mongodb_play/internal/domain/usecases/add_car"
	mongodb "mongodb_play/internal/drivers/mongo_db"
	"mongodb_play/internal/infrastructure/repo"
	"mongodb_play/tests/suits"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddCar_HappyPass(t *testing.T) {
	var c entities.Car
	plates := []string{"E456TT777", "O876CC198", "X261OP159", "A123BE78"}
	fakeCar := gofakeit.Car()
	c.Brand = fakeCar.Brand
	c.Model = fakeCar.Model
	c.Engine = entities.Diesel
	c.EnginePower = 100 * rand.Float32()
	c.NumberPlate = gofakeit.RandomString(plates)

	ctx, s := suits.New(t)

	mongoConnStr := mongodb.MakeConnString(
		s.Config.MongoDB.Username,
		s.Config.MongoDB.Password,
		s.Config.MongoDB.Host,
		s.Config.MongoDB.Port,
	)
	mongoClient, err := mongodb.New(ctx, mongoConnStr)
	if err != nil {
		t.Error("can't connect to mongo db")
	}
	mongoRepo := repo.NewRepo(
		mongoClient,
		s.Config.MongoDB.Database,
		s.Config.MongoDB.Collection,
	)

	add := addcar.New(mongoRepo)

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
	var c entities.Car
	plates := []string{"E456TT777", "O876CC198", "X261OP159", "A123BE78"}
	fakeCar := gofakeit.Car()
	c.Brand = fakeCar.Brand
	c.Model = fakeCar.Model
	c.Engine = entities.Diesel
	c.EnginePower = 100 * rand.Float32()
	c.NumberPlate = gofakeit.RandomString(plates)

	ctx, s := suits.New(t)

	mongoConnStr := mongodb.MakeConnString(
		s.Config.MongoDB.Username,
		s.Config.MongoDB.Password,
		s.Config.MongoDB.Host,
		s.Config.MongoDB.Port,
	)
	mongoClient, err := mongodb.New(ctx, mongoConnStr)
	if err != nil {
		t.Error("can't connect to mongo db")
	}
	mongoRepo := repo.NewRepo(
		mongoClient,
		s.Config.MongoDB.Database,
		s.Config.MongoDB.Collection,
	)

	add := addcar.New(mongoRepo)

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
