package entities

import "github.com/pkg/errors"

type EngineType string

const (
	Gasoline = EngineType("gasoline engine")
	Diesel   = EngineType("diesel engine")
	EV       = EngineType("electrical vehicle")
)

type Car struct {
	Brand       string     `bson:"brand"`
	Model       string     `bson:"model"`
	Engine      EngineType `bson:"engine"`
	EnginePower float32    `bson:"engine_power"`
	NumberPlate string     `bson:"number_plate"`
}

func ValidateCar(c Car) error {
	if c.Brand == "" {
		return errors.New("empty brand")
	}

	if c.Model == "" {
		return errors.New("empty model")
	}

	if string(c.Engine) == "" {
		return errors.New("empty engine")
	}

	if c.EnginePower == 0 {
		return errors.New("empty engine power")
	}

	if c.NumberPlate == "" {
		return errors.New("empty number plate")
	}

	return nil
}
