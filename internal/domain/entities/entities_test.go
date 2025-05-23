package entities

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testInput  = "test"
	emptyInput = ""
)

var (
	carWOBrand = Car{
		Brand:       emptyInput,
		Model:       testInput,
		Engine:      Gasoline,
		EnginePower: 100,
		NumberPlate: testInput,
	}
	carWOModel = Car{
		Brand:       testInput,
		Model:       emptyInput,
		Engine:      Gasoline,
		EnginePower: 100,
		NumberPlate: testInput,
	}
	carWOEngine = Car{
		Brand:       testInput,
		Model:       testInput,
		Engine:      EngineType(10),
		EnginePower: 100,
		NumberPlate: testInput,
	}
	carWOEnginePower = Car{
		Brand:       testInput,
		Model:       testInput,
		Engine:      Diesel,
		EnginePower: 0,
		NumberPlate: testInput,
	}
	carWONumberPlate = Car{
		Brand:       testInput,
		Model:       testInput,
		Engine:      Gasoline,
		EnginePower: 100,
		NumberPlate: emptyInput,
	}
	carWOBrandAndModel = Car{
		Brand:       emptyInput,
		Model:       emptyInput,
		Engine:      Gasoline,
		EnginePower: 100,
		NumberPlate: testInput,
	}
	carWOBrandAndModelAndNumberPlate = Car{
		Brand:       emptyInput,
		Model:       emptyInput,
		Engine:      Gasoline,
		EnginePower: 100,
		NumberPlate: emptyInput,
	}
	carWithAllFieldsCompleted = Car{
		Brand:       testInput,
		Model:       testInput,
		Engine:      Gasoline,
		EnginePower: 100,
		NumberPlate: testInput,
	}
	// test table.
	inputs = []struct {
		name        string
		input       Car
		expectedErr error
	}{
		{name: "car w/o Brand Info", input: carWOBrand, expectedErr: ErrEmptyBrand},
		{name: "car w/o Model Info", input: carWOModel, expectedErr: ErrEmptyModel},
		{name: "car w/o Engine Info", input: carWOEngine, expectedErr: ErrEmptyEngine},
		{name: "car w/o Engine Power Info", input: carWOEnginePower, expectedErr: ErrEmptyEnginePower},
		{name: "car w/o Number Plate Info", input: carWONumberPlate, expectedErr: ErrEmptyNumberPlate},
		{
			name:        "car w/o Brand and Model Info",
			input:       carWOBrandAndModel,
			expectedErr: errors.Wrap(ErrEmptyBrand, ErrEmptyModel.Error()),
		},
		{
			name:        "car w/o Brand, Model and Number Plate Info",
			input:       carWOBrandAndModelAndNumberPlate,
			expectedErr: errors.Wrap(errors.Wrap(ErrEmptyBrand, ErrEmptyModel.Error()), ErrEmptyNumberPlate.Error()),
		},
		{name: "car w/ All fields completed", input: carWithAllFieldsCompleted, expectedErr: nil},
	}
)

func Test_ValidateCar(t *testing.T) {
	for _, test := range inputs {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateCar(test.input)
			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_ErrWrapper(t *testing.T) {
	const (
		testErrMsg1 = "test error msg 1"
		testErrMsg2 = "test error msg 2"
	)

	err := ErrWrapper(nil, testErrMsg1)
	require.EqualError(t, err, testErrMsg1)

	err2 := ErrWrapper(err, testErrMsg2)
	assert.EqualError(t, err2, errors.Wrap(err, testErrMsg2).Error())
}
