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

var (
	ErrEmptyBrand       = errors.New("empty brand")
	ErrEmptyModel       = errors.New("empty model")
	ErrEmptyEngine      = errors.New("empty engine")
	ErrEmptyEnginePower = errors.New("empty engine power")
	ErrEmptyNumberPlate = errors.New("empty number plate")
)

func ValidateCar(c Car) error {
	var err error

	if c.Brand == "" {
		err = ErrWrapper(err, ErrEmptyBrand.Error())
	}

	if c.Model == "" {
		err = ErrWrapper(err, ErrEmptyModel.Error())
	}

	if _, ok := EngineTypes[c.Engine]; !ok {
		err = ErrWrapper(err, ErrEmptyEngine.Error())
	}

	if c.EnginePower == 0 {
		err = ErrWrapper(err, ErrEmptyEnginePower.Error())
	}

	if c.NumberPlate == "" {
		err = ErrWrapper(err, ErrEmptyNumberPlate.Error())
	}

	return err
}

func ErrWrapper(err error, msg string) error {
	if err != nil {
		return errors.Wrap(err, msg)
	}
	return errors.New(msg)
}
