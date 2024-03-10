package main

var patterns = []string{"Blinker", "Toad", "Pulsar", "Glider", "Heavy- weight spaceship", "Gosper glider gun"}

var drawPatternCallback = map[string]func(b *board, midX, midY int){
	"Blinker": func(b *board, midX, midY int) {
		b.setCell(midX-1, midY)
		b.setCell(midX, midY)
		b.setCell(midX+1, midY)
	},
	"Toad": func(b *board, midX, midY int) {
		b.setCell(midX-1, midY)
		b.setCell(midX, midY)
		b.setCell(midX+1, midY)

		b.setCell(midX-2, midY+1)
		b.setCell(midX-1, midY+1)
		b.setCell(midX, midY+1)
	},
	"Pulsar": func(b *board, midX, midY int) {
		b.setCell(midX-2, midY-6)
		b.setCell(midX-3, midY-6)
		b.setCell(midX-4, midY-6)

		b.setCell(midX+2, midY-6)
		b.setCell(midX+3, midY-6)
		b.setCell(midX+4, midY-6)

		b.setCell(midX-6, midY-4)
		b.setCell(midX-6, midY-3)
		b.setCell(midX-6, midY-2)

		b.setCell(midX-1, midY-4)
		b.setCell(midX-1, midY-3)
		b.setCell(midX-1, midY-2)

		b.setCell(midX+1, midY-4)
		b.setCell(midX+1, midY-3)
		b.setCell(midX+1, midY-2)

		b.setCell(midX+6, midY-4)
		b.setCell(midX+6, midY-3)
		b.setCell(midX+6, midY-2)

		b.setCell(midX-2, midY-1)
		b.setCell(midX-3, midY-1)
		b.setCell(midX-4, midY-1)

		b.setCell(midX+2, midY-1)
		b.setCell(midX+3, midY-1)
		b.setCell(midX+4, midY-1)

		b.setCell(midX-2, midY+1)
		b.setCell(midX-3, midY+1)
		b.setCell(midX-4, midY+1)

		b.setCell(midX+2, midY+1)
		b.setCell(midX+3, midY+1)
		b.setCell(midX+4, midY+1)

		b.setCell(midX-6, midY+2)
		b.setCell(midX-6, midY+3)
		b.setCell(midX-6, midY+4)

		b.setCell(midX-1, midY+2)
		b.setCell(midX-1, midY+3)
		b.setCell(midX-1, midY+4)

		b.setCell(midX+1, midY+2)
		b.setCell(midX+1, midY+3)
		b.setCell(midX+1, midY+4)

		b.setCell(midX+6, midY+2)
		b.setCell(midX+6, midY+3)
		b.setCell(midX+6, midY+4)

		b.setCell(midX-2, midY+6)
		b.setCell(midX-3, midY+6)
		b.setCell(midX-4, midY+6)

		b.setCell(midX+2, midY+6)
		b.setCell(midX+3, midY+6)
		b.setCell(midX+4, midY+6)
	},
	"Glider": func(b *board, midX, midY int) {
		b.setCell(midX-1, midY+1)
		b.setCell(midX, midY+1)
		b.setCell(midX+1, midY+1)

		b.setCell(midX+1, midY)
		b.setCell(midX, midY-1)
	},
	"Heavy- weight spaceship": func(b *board, midX, midY int) {
		b.setCell(midX-3, midY)
		b.setCell(midX-2, midY)
		b.setCell(midX-1, midY)
		b.setCell(midX, midY)
		b.setCell(midX+1, midY)
		b.setCell(midX+2, midY)

		b.setCell(midX-3, midY-1)
		b.setCell(midX-3, midY-2)
		b.setCell(midX-2, midY-3)
		b.setCell(midX, midY-4)
		b.setCell(midX+1, midY-4)
		b.setCell(midX+3, midY-3)
		b.setCell(midX+3, midY-1)
	},
	"Gosper glider gun": func(b *board, midX, midY int) {
		b.setCell(midX, midY)
		b.setCell(midX-1, midY-1)
		b.setCell(midX-1, midY)
		b.setCell(midX-1, midY+1)
		b.setCell(midX-2, midY-2)
		b.setCell(midX-2, midY+2)
		b.setCell(midX-3, midY)
		b.setCell(midX-4, midY-3)
		b.setCell(midX-4, midY+3)
		b.setCell(midX-5, midY-3)
		b.setCell(midX-5, midY+3)
		b.setCell(midX-6, midY-2)
		b.setCell(midX-6, midY+2)
		b.setCell(midX-7, midY-1)
		b.setCell(midX-7, midY)
		b.setCell(midX-7, midY+1)

		b.setCell(midX-16, midY-1)
		b.setCell(midX-16, midY)
		b.setCell(midX-17, midY-1)
		b.setCell(midX-17, midY)

		b.setCell(midX+3, midY-3)
		b.setCell(midX+3, midY-2)
		b.setCell(midX+3, midY-1)
		b.setCell(midX+4, midY-3)
		b.setCell(midX+4, midY-2)
		b.setCell(midX+4, midY-1)
		b.setCell(midX+5, midY-4)
		b.setCell(midX+5, midY)
		b.setCell(midX+7, midY-4)
		b.setCell(midX+7, midY)
		b.setCell(midX+7, midY-5)
		b.setCell(midX+7, midY+1)

		b.setCell(midX+17, midY-2)
		b.setCell(midX+17, midY-3)
		b.setCell(midX+18, midY-2)
		b.setCell(midX+18, midY-3)
	},
}
