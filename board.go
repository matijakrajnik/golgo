package main

const (
	boardWidth  = 60
	boardHeight = 40
)

type board struct {
	genStart             [][]bool
	genCurrent           [][]bool
	genNext              [][]bool
	generation           int
	width, height        int
	xCellSize, yCellSize int
	infinite             bool
}

func newBoard(width, height int) *board {
	board := &board{
		width:      width,
		height:     height,
		generation: 0,
	}
	board.initGrid()

	return board
}

func (b *board) restart() {
	for i := 0; i < b.height; i++ {
		copy(b.genCurrent[i], b.genStart[i])
	}
}

func (b *board) initGrid() {
	b.genStart = make([][]bool, b.height)
	b.genCurrent = make([][]bool, b.height)
	b.genNext = make([][]bool, b.height)

	for i := 0; i < b.height; i++ {
		b.genStart[i] = make([]bool, b.width)
		b.genCurrent[i] = make([]bool, b.width)
		b.genNext[i] = make([]bool, b.width)
	}
}

func (b *board) nextGen() {
	if b.generation == 0 {
		b.saveStartPattern()
	}

	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			n := b.countNeighbours(x, y)
			if b.genCurrent[y][x] {
				if n < 2 {
					b.genNext[y][x] = false
				} else if n == 2 || n == 3 {
					b.genNext[y][x] = true
				} else if n > 3 {
					b.genNext[y][x] = false
				}
			} else {
				if n == 3 {
					b.genNext[y][x] = true
				} else {
					b.genNext[y][x] = false
				}
			}
		}
	}

	b.generation++
	b.genCurrent, b.genNext = b.genNext, b.genCurrent
}

func (b *board) saveStartPattern() {
	for i := 0; i < b.height; i++ {
		copy(b.genStart[i], b.genCurrent[i])
	}
}

func (b *board) isAlive(x, y int) bool {
	if !b.infinite && b.isOverflow(x, y) {
		return false
	}

	if x == -1 {
		x = b.width - 1
	} else if x == b.width {
		x = 0
	}

	if y == -1 {
		y = b.height - 1
	} else if y == b.height {
		y = 0
	}

	return b.genCurrent[y][x]
}

func (b *board) isOverflow(x, y int) bool {
	return x <= -1 || x >= b.width || y <= -1 || y >= b.height
}

func (b *board) countNeighbours(x, y int) int {
	n := 0

	if b.isAlive(x-1, y-1) {
		n++
	}
	if b.isAlive(x, y-1) {
		n++
	}
	if b.isAlive(x+1, y-1) {
		n++
	}
	if b.isAlive(x-1, y) {
		n++
	}
	if b.isAlive(x+1, y) {
		n++
	}
	if b.isAlive(x-1, y+1) {
		n++
	}
	if b.isAlive(x, y+1) {
		n++
	}
	if b.isAlive(x+1, y+1) {
		n++
	}

	return n
}

func (b *board) calculateCellSize(w, h int) {
	b.xCellSize = w / b.width
	b.yCellSize = h / b.height
	if b.xCellSize < b.yCellSize {
		b.yCellSize = b.xCellSize
	} else {
		b.xCellSize = b.yCellSize
	}
}

func (b *board) calculateOffset(w, h int) (int, int) {
	return (w - b.width*b.xCellSize) / 2, (h - b.height*b.yCellSize) / 2
}
