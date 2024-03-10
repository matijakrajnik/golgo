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
			g.board.setCell(midX-1, midY)
			g.board.setCell(midX, midY)
			g.board.setCell(midX+1, midY)
		case "Toad":
			g.board.setCell(midX-1, midY)
			g.board.setCell(midX, midY)
			g.board.setCell(midX+1, midY)

			g.board.setCell(midX-2, midY+1)
			g.board.setCell(midX-1, midY+1)
			g.board.setCell(midX, midY+1)
		case "Pulsar":
			g.board.setCell(midX-2, midY-6)
			g.board.setCell(midX-3, midY-6)
			g.board.setCell(midX-4, midY-6)

			g.board.setCell(midX+2, midY-6)
			g.board.setCell(midX+3, midY-6)
			g.board.setCell(midX+4, midY-6)

			g.board.setCell(midX-6, midY-4)
			g.board.setCell(midX-6, midY-3)
			g.board.setCell(midX-6, midY-2)

			g.board.setCell(midX-1, midY-4)
			g.board.setCell(midX-1, midY-3)
			g.board.setCell(midX-1, midY-2)

			g.board.setCell(midX+1, midY-4)
			g.board.setCell(midX+1, midY-3)
			g.board.setCell(midX+1, midY-2)

			g.board.setCell(midX+6, midY-4)
			g.board.setCell(midX+6, midY-3)
			g.board.setCell(midX+6, midY-2)

			g.board.setCell(midX-2, midY-1)
			g.board.setCell(midX-3, midY-1)
			g.board.setCell(midX-4, midY-1)

			g.board.setCell(midX+2, midY-1)
			g.board.setCell(midX+3, midY-1)
			g.board.setCell(midX+4, midY-1)

			g.board.setCell(midX-2, midY+1)
			g.board.setCell(midX-3, midY+1)
			g.board.setCell(midX-4, midY+1)

			g.board.setCell(midX+2, midY+1)
			g.board.setCell(midX+3, midY+1)
			g.board.setCell(midX+4, midY+1)

			g.board.setCell(midX-6, midY+2)
			g.board.setCell(midX-6, midY+3)
			g.board.setCell(midX-6, midY+4)

			g.board.setCell(midX-1, midY+2)
			g.board.setCell(midX-1, midY+3)
			g.board.setCell(midX-1, midY+4)

			g.board.setCell(midX+1, midY+2)
			g.board.setCell(midX+1, midY+3)
			g.board.setCell(midX+1, midY+4)

			g.board.setCell(midX+6, midY+2)
			g.board.setCell(midX+6, midY+3)
			g.board.setCell(midX+6, midY+4)

			g.board.setCell(midX-2, midY+6)
			g.board.setCell(midX-3, midY+6)
			g.board.setCell(midX-4, midY+6)

			g.board.setCell(midX+2, midY+6)
			g.board.setCell(midX+3, midY+6)
			g.board.setCell(midX+4, midY+6)
		case "Glider":
			g.board.setCell(midX-1, midY+1)
			g.board.setCell(midX, midY+1)
			g.board.setCell(midX+1, midY+1)

			g.board.setCell(midX+1, midY)
			g.board.setCell(midX, midY-1)
		case "Heavy- weight spaceship":
			g.board.setCell(midX-3, midY)
			g.board.setCell(midX-2, midY)
			g.board.setCell(midX-1, midY)
			g.board.setCell(midX, midY)
			g.board.setCell(midX+1, midY)
			g.board.setCell(midX+2, midY)

			g.board.setCell(midX-3, midY-1)
			g.board.setCell(midX-3, midY-2)
			g.board.setCell(midX-2, midY-3)
			g.board.setCell(midX, midY-4)
			g.board.setCell(midX+1, midY-4)
			g.board.setCell(midX+3, midY-3)
			g.board.setCell(midX+3, midY-1)
		case "Gosper glider gun":
			g.board.setCell(midX, midY)
			g.board.setCell(midX-1, midY-1)
			g.board.setCell(midX-1, midY)
			g.board.setCell(midX-1, midY+1)
			g.board.setCell(midX-2, midY-2)
			g.board.setCell(midX-2, midY+2)
			g.board.setCell(midX-3, midY)
			g.board.setCell(midX-4, midY-3)
			g.board.setCell(midX-4, midY+3)
			g.board.setCell(midX-5, midY-3)
			g.board.setCell(midX-5, midY+3)
			g.board.setCell(midX-6, midY-2)
			g.board.setCell(midX-6, midY+2)
			g.board.setCell(midX-7, midY-1)
			g.board.setCell(midX-7, midY)
			g.board.setCell(midX-7, midY+1)

			g.board.setCell(midX-16, midY-1)
			g.board.setCell(midX-16, midY)
			g.board.setCell(midX-17, midY-1)
			g.board.setCell(midX-17, midY)

			g.board.setCell(midX+3, midY-3)
			g.board.setCell(midX+3, midY-2)
			g.board.setCell(midX+3, midY-1)
			g.board.setCell(midX+4, midY-3)
			g.board.setCell(midX+4, midY-2)
			g.board.setCell(midX+4, midY-1)
			g.board.setCell(midX+5, midY-4)
			g.board.setCell(midX+5, midY)
			g.board.setCell(midX+7, midY-4)
			g.board.setCell(midX+7, midY)
			g.board.setCell(midX+7, midY-5)
			g.board.setCell(midX+7, midY+1)

			g.board.setCell(midX+17, midY-2)
			g.board.setCell(midX+17, midY-3)
			g.board.setCell(midX+18, midY-2)
			g.board.setCell(midX+18, midY-3)
		}
		g.reset()
	})
}
