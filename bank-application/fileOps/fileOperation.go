package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	defaultValue := 1000.0
	if err != nil {

		return defaultValue, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}
	return value, nil
}
