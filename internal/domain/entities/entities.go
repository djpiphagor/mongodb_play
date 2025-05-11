package entities

import "github.com/pkg/errors"

type EngineType uint8

const (
	UndefinedType EngineType = iota
	Gasoline
	Diesel
	EV
)

var EngineTypes = map[string]EngineType{
	"gasoline engine":  Gasoline,
	"diesel engine":    Diesel,
	"electric vehicle": EV,
}

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

	if c.Engine != UndefinedType {
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
