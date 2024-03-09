package main

import (
	"fyne.io/fyne/v2/widget"
)

var patterns = []string{"Blinker", "Toad", "Pulsar", "Glider", "Heavy- weight spaceship", "Gosper glider gun"}

func generatePatterns(g *game) *widget.Select {
	return widget.NewSelect(patterns, func(pattern string) {
		g.board.initGrid()
		midX, midY := g.board.width/2, g.board.height/2
		switch pattern {
		case "Blinker":
			g.board.genCurrent[midY][midX-1] = true
			g.board.genCurrent[midY][midX] = true
			g.board.genCurrent[midY][midX+1] = true
		case "Toad":
			g.board.genCurrent[midY][midX-1] = true
			g.board.genCurrent[midY][midX] = true
			g.board.genCurrent[midY][midX+1] = true

			g.board.genCurrent[midY+1][midX-2] = true
			g.board.genCurrent[midY+1][midX-1] = true
			g.board.genCurrent[midY+1][midX] = true
		case "Pulsar":
			g.board.genCurrent[midY-6][midX-2] = true
			g.board.genCurrent[midY-6][midX-3] = true
			g.board.genCurrent[midY-6][midX-4] = true

			g.board.genCurrent[midY-6][midX+2] = true
			g.board.genCurrent[midY-6][midX+3] = true
			g.board.genCurrent[midY-6][midX+4] = true

			g.board.genCurrent[midY-4][midX-6] = true
			g.board.genCurrent[midY-3][midX-6] = true
			g.board.genCurrent[midY-2][midX-6] = true

			g.board.genCurrent[midY-4][midX-1] = true
			g.board.genCurrent[midY-3][midX-1] = true
			g.board.genCurrent[midY-2][midX-1] = true

			g.board.genCurrent[midY-4][midX+1] = true
			g.board.genCurrent[midY-3][midX+1] = true
			g.board.genCurrent[midY-2][midX+1] = true

			g.board.genCurrent[midY-4][midX+6] = true
			g.board.genCurrent[midY-3][midX+6] = true
			g.board.genCurrent[midY-2][midX+6] = true

			g.board.genCurrent[midY-1][midX-2] = true
			g.board.genCurrent[midY-1][midX-3] = true
			g.board.genCurrent[midY-1][midX-4] = true

			g.board.genCurrent[midY-1][midX+2] = true
			g.board.genCurrent[midY-1][midX+3] = true
			g.board.genCurrent[midY-1][midX+4] = true

			g.board.genCurrent[midY+1][midX-2] = true
			g.board.genCurrent[midY+1][midX-3] = true
			g.board.genCurrent[midY+1][midX-4] = true

			g.board.genCurrent[midY+1][midX+2] = true
			g.board.genCurrent[midY+1][midX+3] = true
			g.board.genCurrent[midY+1][midX+4] = true

			g.board.genCurrent[midY+2][midX-6] = true
			g.board.genCurrent[midY+3][midX-6] = true
			g.board.genCurrent[midY+4][midX-6] = true

			g.board.genCurrent[midY+2][midX-1] = true
			g.board.genCurrent[midY+3][midX-1] = true
			g.board.genCurrent[midY+4][midX-1] = true

			g.board.genCurrent[midY+2][midX+1] = true
			g.board.genCurrent[midY+3][midX+1] = true
			g.board.genCurrent[midY+4][midX+1] = true

			g.board.genCurrent[midY+2][midX+6] = true
			g.board.genCurrent[midY+3][midX+6] = true
			g.board.genCurrent[midY+4][midX+6] = true

			g.board.genCurrent[midY+6][midX-2] = true
			g.board.genCurrent[midY+6][midX-3] = true
			g.board.genCurrent[midY+6][midX-4] = true

			g.board.genCurrent[midY+6][midX+2] = true
			g.board.genCurrent[midY+6][midX+3] = true
			g.board.genCurrent[midY+6][midX+4] = true
		case "Glider":
			g.board.genCurrent[midY+1][midX-1] = true
			g.board.genCurrent[midY+1][midX] = true
			g.board.genCurrent[midY+1][midX+1] = true

			g.board.genCurrent[midY][midX+1] = true
			g.board.genCurrent[midY-1][midX] = true
		case "Heavy- weight spaceship":
			g.board.genCurrent[midY][midX-3] = true
			g.board.genCurrent[midY][midX-2] = true
			g.board.genCurrent[midY][midX-1] = true
			g.board.genCurrent[midY][midX] = true
			g.board.genCurrent[midY][midX+1] = true
			g.board.genCurrent[midY][midX+2] = true

			g.board.genCurrent[midY-1][midX-3] = true
			g.board.genCurrent[midY-2][midX-3] = true
			g.board.genCurrent[midY-3][midX-2] = true
			g.board.genCurrent[midY-4][midX] = true
			g.board.genCurrent[midY-4][midX+1] = true
			g.board.genCurrent[midY-3][midX+3] = true
			g.board.genCurrent[midY-1][midX+3] = true
		case "Gosper glider gun":
			g.board.genCurrent[midY][midX] = true
			g.board.genCurrent[midY-1][midX-1] = true
			g.board.genCurrent[midY][midX-1] = true
			g.board.genCurrent[midY+1][midX-1] = true
			g.board.genCurrent[midY-2][midX-2] = true
			g.board.genCurrent[midY+2][midX-2] = true
			g.board.genCurrent[midY][midX-3] = true
			g.board.genCurrent[midY-3][midX-4] = true
			g.board.genCurrent[midY+3][midX-4] = true
			g.board.genCurrent[midY-3][midX-5] = true
			g.board.genCurrent[midY+3][midX-5] = true
			g.board.genCurrent[midY-2][midX-6] = true
			g.board.genCurrent[midY+2][midX-6] = true
			g.board.genCurrent[midY-1][midX-7] = true
			g.board.genCurrent[midY][midX-7] = true
			g.board.genCurrent[midY+1][midX-7] = true

			g.board.genCurrent[midY-1][midX-16] = true
			g.board.genCurrent[midY][midX-16] = true
			g.board.genCurrent[midY-1][midX-17] = true
			g.board.genCurrent[midY][midX-17] = true

			g.board.genCurrent[midY-3][midX+3] = true
			g.board.genCurrent[midY-2][midX+3] = true
			g.board.genCurrent[midY-1][midX+3] = true
			g.board.genCurrent[midY-3][midX+4] = true
			g.board.genCurrent[midY-2][midX+4] = true
			g.board.genCurrent[midY-1][midX+4] = true
			g.board.genCurrent[midY-4][midX+5] = true
			g.board.genCurrent[midY][midX+5] = true
			g.board.genCurrent[midY-4][midX+7] = true
			g.board.genCurrent[midY][midX+7] = true
			g.board.genCurrent[midY-5][midX+7] = true
			g.board.genCurrent[midY+1][midX+7] = true

			g.board.genCurrent[midY-2][midX+17] = true
			g.board.genCurrent[midY-3][midX+17] = true
			g.board.genCurrent[midY-2][midX+18] = true
			g.board.genCurrent[midY-3][midX+18] = true
		}
		g.reset()
	})
}
