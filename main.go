package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Game of life")
	game := newGame()
	window.SetContent(game.buildUI())
	window.Resize(fyne.Size{Width: 900, Height: 600})
	window.CenterOnScreen()
	game.run()
	window.ShowAndRun()
}
