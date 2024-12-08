package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloattFromFile(filename string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func WriteFloatToFile(filename string, value float64) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(balanceText), 0o644)
}
