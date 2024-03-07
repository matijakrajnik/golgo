package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const appID = "com.github.matijakrajnik.golgo"

func main() {
	app := app.NewWithID(appID)
	loadPreferences(app)
	window := app.NewWindow("Game of life")
	game := newGame()
	window.SetContent(game.buildUI())
	window.Resize(fyne.Size{Width: 900, Height: 600})
	window.CenterOnScreen()
	game.run()
	window.ShowAndRun()
}
