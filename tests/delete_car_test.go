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

	addCarUC := addcar.New(mongoRepo)
	_, err := addCarUC.Action(ctx, s.Config.MongoDB.Collection, testCar)
	require.NoError(t, err)

	deleteCarUC := deletecar.New(mongoRepo)
	err = deleteCarUC.Action(ctx, s.Config.MongoDB.Collection, testCar.NumberPlate)
	require.NoError(t, err)

	_, err = mongoRepo.FindCarByNumberPlate(ctx, testCar.NumberPlate)
	require.ErrorIs(t, mongo.ErrNoDocuments, err)
}
