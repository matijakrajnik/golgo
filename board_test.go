package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFailCoordMsg = func(x, y int) string { return fmt.Sprintf("Unexpected value at [%d][%d]", x, y) }

func TestBoardNextGen(t *testing.T) {
	t.Run("Infinite", func(t *testing.T) {
		b := newBoard(4, 4, true)

		b.setStartingPattern([][]bool{
			{true, true, false, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		})

		b.nextGen()

		expected := [][]bool{
			{true, false, false, false},
			{true, false, false, false},
			{false, false, false, false},
			{true, false, false, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("NotInfinite", func(t *testing.T) {
		b := newBoard(4, 4, false)

		b.setStartingPattern([][]bool{
			{true, true, false, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		})

		b.nextGen()

		expected := [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("AliveCellWithoutAliveNeighbourDies", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		})

		b.nextGen()

		expected := [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("AliveCellWithOneAliveNeighbourDies", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{false, true, false},
			{false, true, false},
			{false, false, false},
		})

		b.nextGen()

		expected := [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("AliveCellWithTwoAliveNeighbourLives", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{false, false, true},
			{false, true, false},
			{true, false, false},
		})

		b.nextGen()

		expected := [][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("AliveCellWithThreeAliveNeighbourLives", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{true, false, true},
			{false, true, false},
			{true, false, false},
		})

		b.nextGen()

		expected := [][]bool{
			{false, true, false},
			{true, true, false},
			{false, false, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("AliveCellWithFourAliveNeighbourDies", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{true, false, true},
			{false, true, false},
			{true, false, true},
		})

		b.nextGen()

		expected := [][]bool{
			{false, true, false},
			{true, false, true},
			{false, true, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("DeadCellWithThreeAliveNeighbourResurects", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{true, false, true},
			{false, true, false},
			{false, false, false},
		})

		b.nextGen()

		expected := [][]bool{
			{false, true, false},
			{false, true, false},
			{false, false, false},
		}

		assert.EqualValues(t, expected, b.genCurrent)
	})
}

func TestSaveStartPattern(t *testing.T) {
	b := newBoard(3, 3, false)

	b.setCell(0, 0)
	b.setCell(1, 0)
	b.setCell(2, 0)

	b.saveStartPattern()

	expected := [][]bool{
		{true, true, true},
		{false, false, false},
		{false, false, false},
	}

	assert.EqualValues(t, expected, b.genCurrent)
}

func TestRestart(t *testing.T) {
	b := newBoard(3, 3, false)

	b.setCell(0, 0)
	b.setCell(1, 0)
	b.setCell(2, 0)

	expected := [][]bool{
		{true, true, true},
		{false, false, false},
		{false, false, false},
	}

	b.nextGen()

	assert.NotEqualValues(t, expected, b.genCurrent)
	b.restart()
	assert.EqualValues(t, expected, b.genCurrent)
}

func TestSetStartingPattern(t *testing.T) {
	b := newBoard(4, 4, false)

	b.setStartingPattern([][]bool{
		{true, true, false, true},
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
	})

	expected := [][]bool{
		{true, true, false, true},
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
	}
	assert.EqualValues(t, expected, b.genCurrent)
}

func TestSetCell(t *testing.T) {
	t.Run("CoordinatesWithinBorder", func(t *testing.T) {
		b := newBoard(2, 2, false)

		b.setCell(1, 0)

		expected := [][]bool{
			{false, true},
			{false, false},
		}
		assert.EqualValues(t, expected, b.genCurrent)
	})

	t.Run("CoordinatesWithoutBorder", func(t *testing.T) {
		b := newBoard(2, 2, false)

		b.setCell(-1, -1)

		expected := [][]bool{
			{false, false},
			{false, false},
		}
		assert.EqualValues(t, expected, b.genCurrent)
	})
}

func TestCalculateCellSize(t *testing.T) {
	b := newBoard(33, 22, false)

	b.calculateCellSize(600, 500)

	assert.EqualValues(t, 18, b.xCellSize)
	assert.EqualValues(t, 18, b.yCellSize)
}

func TestCalculateOffset(t *testing.T) {
	t.Run("BiggerWidth", func(t *testing.T) {
		b := newBoard(33, 22, false)
		b.calculateCellSize(600, 500)

		offsetX, offsetY := b.calculateOffset(600, 500)
		assert.EqualValues(t, 3, offsetX)
		assert.EqualValues(t, 52, offsetY)
	})

	t.Run("BiggerHeight", func(t *testing.T) {
		b := newBoard(22, 33, false)
		b.calculateCellSize(500, 600)

		offsetX, offsetY := b.calculateOffset(500, 600)
		assert.EqualValues(t, 52, offsetX)
		assert.EqualValues(t, 3, offsetY)
	})
}

func BenchmarkNewBoard(b *testing.B) {
	sizes := []int{10, 50, 100}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("BoardSize: %d", size), func(b *testing.B) {
			newBoard(size, size, true)
		})
	}
}

func BenchmarkBoardNextGen(b *testing.B) {
	sizes := []int{10, 50, 100}

	for _, size := range sizes {
		board := newBoard(size, size, true)
		board.setStartingPattern(createTestImportPattern(size, size))
		b.Run(fmt.Sprintf("BoardSize:%d", size), func(b *testing.B) {
			board.nextGen()
		})
	}
}

func BenchmarkBoardSaveStartPattern(b *testing.B) {
	sizes := []int{10, 50, 100}

	for _, size := range sizes {
		board := newBoard(size, size, true)
		board.setStartingPattern(createTestImportPattern(size, size))
		b.Run(fmt.Sprintf("BoardSize:%d", size), func(b *testing.B) {
			board.saveStartPattern()
		})
	}
}

func BenchmarkBoardSetStartingPattern(b *testing.B) {
	sizes := []int{10, 50, 100}

	for _, size := range sizes {
		board := newBoard(size, size, true)
		b.Run(fmt.Sprintf("BoardSize:%d", size), func(b *testing.B) {
			board.setStartingPattern(createTestImportPattern(size, size))
		})
	}
}
