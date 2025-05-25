package tests_test

import (
	"mongodb_play/internal/domain/entities"
	addcar "mongodb_play/internal/domain/usecases/add_car"
	"mongodb_play/internal/infrastructure/repo"
	"mongodb_play/tests/suits"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindCarByNumberPlate(t *testing.T) {
	ctx, s := suits.New(t)
	mongoRepo := repo.NewRepo(s.Client, s.Config.MongoDB.Database)
	testCar := suits.GeterateRandomCar(t)

	addUC := addcar.New(mongoRepo)
	_, err := addUC.Action(ctx, s.Config.MongoDB.Collection, testCar)
	require.NoError(t, err)

	foundCar, err2 := mongoRepo.FindCarByNumberPlate(ctx, testCar.NumberPlate)
	require.NoError(t, err2)
	assert.Equal(t, testCar, foundCar)

	notExistingNumberPlate := "test"
	foundCar, err2 = mongoRepo.FindCarByNumberPlate(ctx, notExistingNumberPlate)
	require.Error(t, err2)
	assert.Equal(t, entities.Car{}, foundCar)
}
