package tests_test

import (
	addcar "mongodb_play/internal/domain/usecases/add_car"
	"mongodb_play/internal/infrastructure/repo"
	"mongodb_play/tests/suits"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddCar_HappyPass(t *testing.T) {
	ctx, s := suits.New(t)

	mongoRepo := repo.NewRepo(
		s.Client,
		s.Config.MongoDB.Database,
	)
	add := addcar.New(mongoRepo)

	c := suits.GeterateRandomCar(t)
	_, err := add.Action(ctx, s.Config.MongoDB.Collection, c)
	require.NoError(t, err)

	foundCar, err := mongoRepo.FindCarByNumberPlate(ctx, c.NumberPlate)
	require.NoError(t, err)
	assert.Equal(t, c.NumberPlate, foundCar.NumberPlate)
}

func TestAdd2TheSameCars_Failed(t *testing.T) {
	ctx, s := suits.New(t)

	mongoRepo := repo.NewRepo(
		s.Client,
		s.Config.MongoDB.Database,
	)
	add := addcar.New(mongoRepo)

	c := suits.GeterateRandomCar(t)
	_, err := add.Action(ctx, s.Config.MongoDB.Collection, c)
	require.NoError(t, err)
	// here trying to add the same car.
	// It has to end up with an err.
	_, err = add.Action(ctx, s.Config.MongoDB.Collection, c)
	require.Error(t, err)
}
