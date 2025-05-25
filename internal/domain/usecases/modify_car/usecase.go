package modifycar

import (
	"context"
	"mongodb_play/internal/domain/entities"

	"github.com/pkg/errors"
)

type garage interface {
	ModifyCar(context.Context, entities.Car) error
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

func (u *Usecase) Action(ctx context.Context, collName string, newCarData entities.Car) error {
	if e := u.garage.SetCollection(ctx, collName); e != nil {
		return errors.Wrap(e, "cant set collection")
	}
	if !validate(newCarData) {
		return errors.New("got incorrect car data")
	}
	return u.garage.ModifyCar(ctx, newCarData)
}

// validate checks if at least one of the car's data fields is not empty.
func validate(c entities.Car) bool {
	if c.Brand != "" {
		return true
	}
	if c.Model != "" {
		return true
	}
	if _, ok := entities.EngineTypes[c.Engine]; ok {
		return true
	}
	if c.EnginePower != 0 {
		return true
	}
	if c.NumberPlate != "" {
		return true
	}
	return false
}
