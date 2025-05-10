package deletecar

import (
	"context"

	"github.com/pkg/errors"
)

type garage interface {
	DeleteCar(context.Context, string) error
}

type Usecase struct {
	garage garage
}

func New(g garage) *Usecase {
	return &Usecase{
		garage: g,
	}
}

func (u *Usecase) Action(ctx context.Context, nPlate string) error {
	if nPlate == "" {
		return errors.New("got empty number plate to del car")
	}
	return u.garage.DeleteCar(ctx, nPlate)
}
