package tests_test

import (
	"mongodb_play/internal/domain/entities"
	addcar "mongodb_play/internal/domain/usecases/add_car"
	modifycar "mongodb_play/internal/domain/usecases/modify_car"
	"mongodb_play/internal/infrastructure/repo"
	"mongodb_play/tests/suits"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModifyCar(t *testing.T) {
	ctx, s := suits.New(t)
	mongoRepo := repo.NewRepo(s.Client, s.Config.MongoDB.Database)
	testCar := suits.GeterateRandomCar(t)
	emptyCar := entities.Car{}
	emptyCar.Engine = entities.EngineType(10)
	// First, we're trying to add a car to Mongo.
	// It has to end up with no error.
	addCarUC := addcar.New(mongoRepo)
	_, err := addCarUC.Action(ctx, s.Config.MongoDB.Collection, testCar)
	require.NoError(t, err)
	// Here, we're trying to modify the car, using the same number plate.
	modifyCarUC := modifycar.New(mongoRepo)
	testCarBrand := "test brand"
	testCar.Brand = testCarBrand
	err = modifyCarUC.Action(ctx, s.Config.MongoDB.Collection, testCar)
	require.NoError(t, err)
	modifiedCar, err2 := mongoRepo.FindCarByNumberPlate(ctx, testCar.NumberPlate)
	require.NoError(t, err2)
	assert.Equal(t, testCarBrand, modifiedCar.Brand)

	err = modifyCarUC.Action(ctx, s.Config.MongoDB.Collection, emptyCar)
	require.Error(t, err)
}
