package deletecar

import (
	"context"

	"github.com/pkg/errors"
)

type garage interface {
	DeleteCar(context.Context, string) error
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

func (u *Usecase) Action(ctx context.Context, collName string, nPlate string) error {
	if e := u.garage.SetCollection(ctx, collName); e != nil {
		return errors.Wrap(e, "cant set collection")
	}
	if nPlate == "" {
		return errors.New("got empty number plate to del car")
	}
	return u.garage.DeleteCar(ctx, nPlate)
}
