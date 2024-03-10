package main

import (
	"fmt"
	"strconv"
)

const (
	minBoardSize = 10
	maxBoardSize = 100
)

func boardSizeValidator(s string) error {
	returnErr := fmt.Errorf("value must be a number between %d and %d", minBoardSize, maxBoardSize)

	n, err := strconv.Atoi(s)
	if err != nil {
		return returnErr
	}

	if n < minBoardSize || n > maxBoardSize {
		return returnErr
	}

	return nil
}
