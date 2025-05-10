package addcar

import (
	"context"
	"mongodb_play/internal/domain/entities"
)

type garage interface {
	AddCar(context.Context, entities.Car) (string, error)
}

type Usecase struct {
	garage garage
}

func New(g garage) *Usecase {
	return &Usecase{
		garage: g,
	}
}

func (u *Usecase) Action(ctx context.Context, c entities.Car) (string, error) {
	if err := entities.ValidateCar(c); err != nil {
		return "", err
	}
	return u.garage.AddCar(ctx, c)
}
