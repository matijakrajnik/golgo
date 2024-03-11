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

		assert.True(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.False(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))
		assert.False(t, b.genCurrent[0][3], testFailCoordMsg(0, 3))

		assert.True(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))
		assert.False(t, b.genCurrent[1][3], testFailCoordMsg(1, 3))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
		assert.False(t, b.genCurrent[2][3], testFailCoordMsg(2, 3))

		assert.True(t, b.genCurrent[3][0], testFailCoordMsg(3, 0))
		assert.False(t, b.genCurrent[3][1], testFailCoordMsg(3, 1))
		assert.False(t, b.genCurrent[3][2], testFailCoordMsg(3, 2))
		assert.False(t, b.genCurrent[3][3], testFailCoordMsg(3, 3))
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

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.False(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))
		assert.False(t, b.genCurrent[0][3], testFailCoordMsg(0, 3))

		assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))
		assert.False(t, b.genCurrent[1][3], testFailCoordMsg(1, 3))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
		assert.False(t, b.genCurrent[2][3], testFailCoordMsg(2, 3))

		assert.False(t, b.genCurrent[3][0], testFailCoordMsg(3, 0))
		assert.False(t, b.genCurrent[3][1], testFailCoordMsg(3, 1))
		assert.False(t, b.genCurrent[3][2], testFailCoordMsg(3, 2))
		assert.False(t, b.genCurrent[3][3], testFailCoordMsg(3, 3))
	})

	t.Run("LiveCellWithoutLiveNeighbourDies", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		})

		b.nextGen()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.False(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))

		assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
	})

	t.Run("LiveCellWithOneLiveNeighbourDies", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{false, true, false},
			{false, true, false},
			{false, false, false},
		})

		b.nextGen()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.False(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))

		assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
	})

	t.Run("LiveCellWithTwoLiveNeighbourLives", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{false, false, true},
			{false, true, false},
			{true, false, false},
		})

		b.nextGen()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.False(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))

		assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.True(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
	})

	t.Run("LiveCellWithThreeLiveNeighbourLives", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{true, false, true},
			{false, true, false},
			{true, false, false},
		})

		b.nextGen()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.True(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))

		assert.True(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.True(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
	})

	t.Run("LiveCellWithFourLiveNeighbourDies", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{true, false, true},
			{false, true, false},
			{true, false, true},
		})

		b.nextGen()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.True(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))

		assert.True(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.True(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.True(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
	})

	t.Run("DeadCellWithThreeLiveNeighbourResurects", func(t *testing.T) {
		b := newBoard(3, 3, false)

		b.setStartingPattern([][]bool{
			{true, false, true},
			{false, true, false},
			{false, false, false},
		})

		b.nextGen()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.True(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))

		assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.True(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
		assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))

		assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
		assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
		assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
	})
}

func TestSaveStartPattern(t *testing.T) {
	b := newBoard(3, 3, false)

	b.setCell(0, 0)
	b.setCell(1, 0)
	b.setCell(2, 0)

	b.saveStartPattern()

	assert.True(t, b.genStart[0][0], testFailCoordMsg(0, 0))
	assert.True(t, b.genStart[0][1], testFailCoordMsg(0, 1))
	assert.True(t, b.genStart[0][2], testFailCoordMsg(0, 2))

	assert.False(t, b.genStart[1][0], testFailCoordMsg(1, 0))
	assert.False(t, b.genStart[1][1], testFailCoordMsg(1, 1))
	assert.False(t, b.genStart[1][2], testFailCoordMsg(1, 2))

	assert.False(t, b.genStart[2][0], testFailCoordMsg(2, 0))
	assert.False(t, b.genStart[2][1], testFailCoordMsg(2, 1))
	assert.False(t, b.genStart[2][2], testFailCoordMsg(2, 2))
}

func TestRestart(t *testing.T) {
	b := newBoard(3, 3, false)

	b.setCell(0, 0)
	b.setCell(1, 0)
	b.setCell(3, 0)

	b.nextGen()
	b.restart()

	assert.True(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
	assert.True(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
	assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))

	assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
	assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
	assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))

	assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
	assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
	assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
}

func TestSetStartingPattern(t *testing.T) {
	b := newBoard(4, 4, false)

	b.setStartingPattern([][]bool{
		{true, true, false, true},
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
	})

	assert.True(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
	assert.True(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
	assert.False(t, b.genCurrent[0][2], testFailCoordMsg(0, 2))
	assert.True(t, b.genCurrent[0][3], testFailCoordMsg(0, 3))

	assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
	assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
	assert.False(t, b.genCurrent[1][2], testFailCoordMsg(1, 2))
	assert.False(t, b.genCurrent[1][3], testFailCoordMsg(1, 3))

	assert.False(t, b.genCurrent[2][0], testFailCoordMsg(2, 0))
	assert.False(t, b.genCurrent[2][1], testFailCoordMsg(2, 1))
	assert.False(t, b.genCurrent[2][2], testFailCoordMsg(2, 2))
	assert.False(t, b.genCurrent[2][3], testFailCoordMsg(2, 3))

	assert.False(t, b.genCurrent[3][0], testFailCoordMsg(3, 0))
	assert.False(t, b.genCurrent[3][1], testFailCoordMsg(3, 1))
	assert.False(t, b.genCurrent[3][2], testFailCoordMsg(3, 2))
	assert.False(t, b.genCurrent[3][3], testFailCoordMsg(3, 3))

	assert.True(t, b.genStart[0][0], testFailCoordMsg(0, 0))
	assert.True(t, b.genStart[0][1], testFailCoordMsg(0, 1))
	assert.False(t, b.genStart[0][2], testFailCoordMsg(0, 2))
	assert.True(t, b.genStart[0][3], testFailCoordMsg(0, 3))

	assert.False(t, b.genStart[1][0], testFailCoordMsg(1, 0))
	assert.False(t, b.genStart[1][1], testFailCoordMsg(1, 1))
	assert.False(t, b.genStart[1][2], testFailCoordMsg(1, 2))
	assert.False(t, b.genStart[1][3], testFailCoordMsg(1, 3))

	assert.False(t, b.genStart[2][0], testFailCoordMsg(2, 0))
	assert.False(t, b.genStart[2][1], testFailCoordMsg(2, 1))
	assert.False(t, b.genStart[2][2], testFailCoordMsg(2, 2))
	assert.False(t, b.genStart[2][3], testFailCoordMsg(2, 3))

	assert.False(t, b.genStart[3][0], testFailCoordMsg(3, 0))
	assert.False(t, b.genStart[3][1], testFailCoordMsg(3, 1))
	assert.False(t, b.genStart[3][2], testFailCoordMsg(3, 2))
	assert.False(t, b.genStart[3][3], testFailCoordMsg(3, 3))
}

func TestSetCell(t *testing.T) {
	t.Run("CoordinatesWithinBorder", func(t *testing.T) {
		b := newBoard(2, 2, false)

		b.setCell(1, 0)
		b.saveStartPattern()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.True(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
	})

	t.Run("CoordinatesWithoutBorder", func(t *testing.T) {
		b := newBoard(2, 2, false)

		b.setCell(-1, -1)
		b.saveStartPattern()

		assert.False(t, b.genCurrent[0][0], testFailCoordMsg(0, 0))
		assert.False(t, b.genCurrent[0][1], testFailCoordMsg(0, 1))
		assert.False(t, b.genCurrent[1][0], testFailCoordMsg(1, 0))
		assert.False(t, b.genCurrent[1][1], testFailCoordMsg(1, 1))
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
