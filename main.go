package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const appID = "com.github.matijakrajnik.golgo"

var (
	mainWindow fyne.Window
	theGame    *game
)

func main() {
	app := app.NewWithID(appID)
	loadPreferences(app)
	mainWindow = app.NewWindow("Game of life")
	theGame = newGame()
	mainWindow.SetContent(theGame.buildUI())
	mainWindow.Resize(fyne.Size{Width: 900, Height: 600})
	mainWindow.CenterOnScreen()
	theGame.setKeyPressListener(mainWindow)
	theGame.run()
	mainWindow.ShowAndRun()
}
