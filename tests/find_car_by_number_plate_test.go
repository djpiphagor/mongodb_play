package tests_test

import (
	"mongodb_play/internal/domain/entities"
	addcar "mongodb_play/internal/domain/usecases/add_car"
	"mongodb_play/internal/infrastructure/repo"
	"mongodb_play/tests/suits"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFindCarByNumberPlate(t *testing.T) {
	ctx, s := suits.New(t)
	mongoRepo := repo.NewRepo(s.Client, s.Config.MongoDB.Database)
	testCar := suits.GeterateRandomCar(t)
	// First, we're trying to add a car.
	// It has to end up with no error.
	addUC := addcar.New(mongoRepo)
	_, err := addUC.Action(ctx, s.Config.MongoDB.Collection, testCar)
	require.NoError(t, err)
	// Then we're trying to find the car we've just added.
	foundCar, err2 := mongoRepo.FindCarByNumberPlate(ctx, testCar.NumberPlate)
	require.NoError(t, err2)
	assert.Equal(t, testCar, foundCar)
	// Here we're trying to find the non-existing car.
	// So, it has to end up with an error.
	notExistingNumberPlate := "test"
	foundCar, err2 = mongoRepo.FindCarByNumberPlate(ctx, notExistingNumberPlate)
	require.Error(t, err2)
	// The error should be like a mongo db no docs err.
	require.ErrorIs(t, mongo.ErrNoDocuments, err2)
	// Here we check if the empty car returned is the same as the empty car from the Entities.
	assert.Equal(t, entities.Car{}, foundCar)
}
