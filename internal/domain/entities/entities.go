package entities

import "github.com/pkg/errors"

type EngineType uint8

const (
	UndefinedType EngineType = iota
	Gasoline
	Diesel
	EV
)

var EngineTypes = map[EngineType]string{
	UndefinedType: "undefined",
	Gasoline:      "gasoline engine",
	Diesel:        "diesel engine",
	EV:            "electric vehicle",
}

type Car struct {
	Brand       string     `bson:"brand"`
	Model       string     `bson:"model"`
	Engine      EngineType `bson:"engine"`
	EnginePower float32    `bson:"engine_power"`
	NumberPlate string     `bson:"number_plate"`
}

func ValidateCar(c Car) error {
	var err error

	if c.Brand == "" {
		err = errors.Wrap(err, "empty brand")
	}

	if c.Model == "" {
		err = errors.Wrap(err, "empty model")
	}

	if _, ok := EngineTypes[c.Engine]; !ok {
		err = errors.Wrap(err, "empty engine")
	}

	if c.EnginePower == 0 {
		err = errors.Wrap(err, "empty engine power")
	}

	if c.NumberPlate == "" {
		err = errors.Wrap(err, "empty number plate")
	}

	return err
}
