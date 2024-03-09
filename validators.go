package main

import (
	"fmt"
	"strconv"
)

func boardSizeValidator(s string) error {
	min, max := 15, 100
	returnErr := fmt.Errorf("value must be a number between %d and %d", min, max)

	n, err := strconv.Atoi(s)
	if err != nil {
		return returnErr
	}

	if n < min || n > max {
		return returnErr
	}

	return nil
}
