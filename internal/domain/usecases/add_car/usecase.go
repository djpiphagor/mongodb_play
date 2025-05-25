package addcar

import (
	"context"
	"mongodb_play/internal/domain/entities"

	"github.com/pkg/errors"
)

type garage interface {
	AddCar(context.Context, entities.Car) (string, error)
	SetCollection(context.Context, string) error
}

type Usecase struct {
	garage garage
}

func New(g garage) *Usecase {
	return &Usecase{
		garage: g,
	}
}

func (u *Usecase) Action(ctx context.Context, collName string, c entities.Car) (string, error) {
	if e := u.garage.SetCollection(ctx, collName); e != nil {
		return "", errors.Wrap(e, "cant set collection")
	}
	if err := entities.ValidateCar(c); err != nil {
		return "", err
	}
	return u.garage.AddCar(ctx, c)
}
