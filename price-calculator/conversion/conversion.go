package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for stringsIdx, stringsVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringsVal, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to flaot")
		}

		floats[stringsIdx] = floatPrice
	}

	return floats, nil
}
