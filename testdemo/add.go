package testdemo

import (
	"errors"
	"strconv"
)

func Add(a int) int {
	return a
}

func I2S(a int) string {
	return strconv.Itoa(a)
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0不能作除数")
	}

	return a / b, nil
}
