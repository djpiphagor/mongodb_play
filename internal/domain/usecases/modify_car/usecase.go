package modifycar

import (
	"context"
	"mongodb_play/internal/domain/entities"
)

type garage interface {
	ModifyCar(context.Context, entities.Car) error
}

type Usecase struct {
	garage garage
}

func New(g garage) *Usecase {
	return &Usecase{
		garage: g,
	}
}

func (u *Usecase) Action(ctx context.Context, newCarData entities.Car) error {
	if err := entities.ValidateCar(newCarData); err != nil {
		return err
	}
	return u.garage.ModifyCar(ctx, newCarData)
}
