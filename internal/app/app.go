package app

import (
	"context"
	"log/slog"
	"mongodb_play/internal/config"
	"mongodb_play/internal/domain/entities"
	addcar "mongodb_play/internal/domain/usecases/add_car"
	deletecar "mongodb_play/internal/domain/usecases/delete_car"
	modifycar "mongodb_play/internal/domain/usecases/modify_car"
	mongodb "mongodb_play/internal/drivers/mongo_db"
	"mongodb_play/internal/infrastructure/repo"

	"github.com/pkg/errors"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (app *App) Run(ctx context.Context) error {
	mongoConn, err := mongodb.New(ctx,
		mongodb.MakeConnString(
			app.cfg.MongoDB.Username,
			app.cfg.MongoDB.Password,
			app.cfg.MongoDB.Host,
			app.cfg.MongoDB.Port))
	if err != nil {
		return errors.Wrap(err, "error mongo db connect")
	}

	repo := repo.NewRepo(mongoConn, app.cfg.MongoDB.Database, app.cfg.MongoDB.Collection)

	add := addcar.New(repo)
	del := deletecar.New(repo)
	modify := modifycar.New(repo)

	car1 := entities.Car{
		Brand:       "VW",
		Model:       "Polo",
		Engine:      entities.Gasoline,
		EnginePower: 100.3,
		NumberPlate: "A001AA198",
	}
	car2 := entities.Car{
		Brand:       "VW",
		Model:       "Caravelle",
		Engine:      entities.Diesel,
		EnginePower: 78.6,
		NumberPlate: "B001BB777",
	}

	id, err := add.Action(ctx, car1)
	if err != nil {
		slog.Error("error adding car1")
	} else {
		slog.With(slog.String("_id", id)).Info("successfully added")
	}

	id, err = add.Action(ctx, car2)
	if err != nil {
		slog.Error("error adding car2")
	} else {
		slog.With(slog.String("_id", id)).Info("successfully added")
	}

	err = del.Action(ctx, "A001AA198")
	if err != nil {
		slog.Error("error deleting A001AA198")
	}

	car3 := entities.Car{
		Brand:       "VW",
		Model:       "Golf",
		Engine:      entities.Gasoline,
		EnginePower: 160.7,
		NumberPlate: "A001AA198",
	}

	err = modify.Action(ctx, car3)
	if err != nil {
		slog.Error("error")
	}

	return nil
}
