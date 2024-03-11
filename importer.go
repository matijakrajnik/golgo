package main

import (
	"fmt"
	"strings"
)

const (
	aliveSymbol     = "X"
	deadSymbol      = "."
	rowSeparator    = "\n"
	columnSeparator = " "
)

func generateTemplateBytes(pattern [][]bool) []byte {
	s := ""
	for _, row := range pattern {
		for _, value := range row {
			if value {
				s += aliveSymbol
			} else {
				s += deadSymbol
			}
			s += columnSeparator
		}
		s = strings.TrimSuffix(s, columnSeparator)
		s += rowSeparator
	}

	return []byte(s)
}

func parseImportedPattern(bytes []byte) ([][]bool, error) {
	pattern := make([][]bool, 0)
	data := string(bytes)

	for strings.HasSuffix(data, rowSeparator) {
		data = strings.TrimSuffix(data, rowSeparator)
	}

	rows := strings.Split(data, rowSeparator)
	rowsN := len(rows)
	if rowsN < minBoardSize || rowsN > maxBoardSize {
		return pattern, fmt.Errorf("number of rows must be between %d and %d", minBoardSize, maxBoardSize)
	}

	columnsN := len(strings.Split(rows[0], columnSeparator))
	if columnsN < minBoardSize || columnsN > maxBoardSize {
		return pattern, fmt.Errorf("number of columns must be between %d and %d", minBoardSize, maxBoardSize)
	}

	pattern = make([][]bool, rowsN)
	for rowI, row := range rows {
		column := strings.Split(row, columnSeparator)
		if len(column) != columnsN {
			return pattern, fmt.Errorf("unexpected number of columns in row %d", rowI)
		}
		pattern[rowI] = make([]bool, columnsN)
		for columnI, value := range column {
			if value != aliveSymbol && value != deadSymbol {
				return pattern, fmt.Errorf(
					"\ninvalid value '%s' at [%d][%d]"+
						"\nvalid values are:\t\t"+
						"\nalive cell:\t%s\t\t"+
						"\ndead cell:\t%s\t\t",
					value, rowI, columnI, aliveSymbol, deadSymbol,
				)
			}
			pattern[rowI][columnI] = value == aliveSymbol
		}
	}

	return pattern, nil
}
