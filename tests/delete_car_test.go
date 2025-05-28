package tests_test

import (
	addcar "mongodb_play/internal/domain/usecases/add_car"
	deletecar "mongodb_play/internal/domain/usecases/delete_car"
	"mongodb_play/internal/infrastructure/repo"
	"mongodb_play/tests/suits"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteCar(t *testing.T) {
	ctx, s := suits.New(t)
	mongoRepo := repo.NewRepo(s.Client, s.Config.MongoDB.Database)
	testCar := suits.GeterateRandomCar(t)
	// Here we're trying to add a car.
	// It has to end up with no error.
	addCarUC := addcar.New(mongoRepo)
	_, err := addCarUC.Action(ctx, s.Config.MongoDB.Collection, testCar)
	require.NoError(t, err)
	// Here we're trying to remove the car added.
	// It has to end up with no error.
	deleteCarUC := deletecar.New(mongoRepo)
	err = deleteCarUC.Action(ctx, s.Config.MongoDB.Collection, testCar.NumberPlate)
	require.NoError(t, err)
	// Here we're doing a double check.
	// We're trying to find the car deleted by number plate.
	// So, if it's removed from DB, it has to end up with an error.
	// And an error should be like a mongo no docs error.
	_, err = mongoRepo.FindCarByNumberPlate(ctx, testCar.NumberPlate)
	require.ErrorIs(t, mongo.ErrNoDocuments, err)
}
