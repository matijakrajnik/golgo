package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestImportPattern(w, h int) [][]bool {
	pattern := make([][]bool, h)

	for y := range pattern {
		pattern[y] = make([]bool, w)
		for x := range pattern[y] {
			pattern[y][x] = (x+y)%2 == 0
		}
	}

	return pattern
}

func createTestImportPatternBytes(w, h int) []byte {
	s := ""

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (x+y)%2 == 0 {
				s += "X "
			} else {
				s += ". "
			}
		}
		s = strings.TrimSuffix(s, " ")
		s += "\n"
	}

	return []byte(s)
}

func TestGenerateTemplateBytes(t *testing.T) {
	bytes := generateTemplateBytes(createTestImportPattern(30, 20))
	assert.EqualValues(t, bytes, createTestImportPatternBytes(30, 20))
}

func TestParseImportedPattern(t *testing.T) {
	t.Run("ValidPattern", func(t *testing.T) {
		parsed, err := parseImportedPattern(createTestImportPatternBytes(30, 20))

		assert.Nil(t, err)
		assert.EqualValues(t, createTestImportPattern(30, 20), parsed)
	})

	t.Run("NotEnoughRows", func(t *testing.T) {
		parsed, err := parseImportedPattern(createTestImportPatternBytes(30, 9))

		assert.EqualValues(t, "number of rows must be between 10 and 100", err.Error())
		assert.EqualValues(t, make([][]bool, 0), parsed)
	})

	t.Run("NotEnoughColumns", func(t *testing.T) {
		parsed, err := parseImportedPattern(createTestImportPatternBytes(9, 20))

		assert.EqualValues(t, "number of columns must be between 10 and 100", err.Error())
		assert.EqualValues(t, make([][]bool, 0), parsed)
	})

	t.Run("TooManyRows", func(t *testing.T) {
		parsed, err := parseImportedPattern(createTestImportPatternBytes(30, 101))

		assert.EqualValues(t, "number of rows must be between 10 and 100", err.Error())
		assert.EqualValues(t, make([][]bool, 0), parsed)
	})

	t.Run("TooManyColumns", func(t *testing.T) {
		parsed, err := parseImportedPattern(createTestImportPatternBytes(101, 20))

		assert.EqualValues(t, "number of columns must be between 10 and 100", err.Error())
		assert.EqualValues(t, make([][]bool, 0), parsed)
	})

	t.Run("TooManyValuesInSingleColumn", func(t *testing.T) {
		parsed, err := parseImportedPattern([]byte(
			"" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X .\n" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n",
		))

		expected := [][]bool{
			{true, false, true, false, true, false, true, false, true, false},
			{false, true, false, true, false, true, false, true, false, true},
			{true, false, true, false, true, false, true, false, true, false},
			{false, true, false, true, false, true, false, true, false, true},
			{true, false, true, false, true, false, true, false, true, false},
			nil, nil, nil, nil, nil,
		}

		assert.EqualValues(t, "unexpected number of columns in row 5", err.Error())
		assert.EqualValues(t, expected, parsed)
	})

	t.Run("InvalidCharacter", func(t *testing.T) {
		parsed, err := parseImportedPattern([]byte(
			"" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n" +
				"X . X . X . X . X .\n" +
				". X . X . F . X . X\n" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n" +
				"X . X . X . X . X .\n" +
				". X . X . X . X . X\n",
		))

		expected := [][]bool{
			{true, false, true, false, true, false, true, false, true, false},
			{false, true, false, true, false, true, false, true, false, true},
			{true, false, true, false, true, false, true, false, true, false},
			{false, true, false, true, false, true, false, true, false, true},
			{true, false, true, false, true, false, true, false, true, false},
			{false, true, false, true, false, false, false, false, false, false},
			nil, nil, nil, nil,
		}

		expectedError := "" +
			"\ninvalid value 'F' at [5][5]" +
			"\nvalid values are:\t\t" +
			"\nalive cell:\tX\t\t" +
			"\ndead cell:\t.\t\t"

		assert.EqualValues(t, expectedError, err.Error())
		assert.EqualValues(t, expected, parsed)
	})
}
