package main

import "fyne.io/fyne/v2"

var preferences fyne.Preferences

type prefKey int

const (
	boardWidthKey prefKey = iota
	boardHeightKey
	infiniteBoardKey
	speedKey
)

var prefKeys = map[prefKey]string{
	boardWidthKey:    "boardWidth",
	boardHeightKey:   "boardHeight",
	infiniteBoardKey: "infiniteBoard",
	speedKey:         "speed",
}

func loadPreferences(app fyne.App) {
	preferences = app.Preferences()
}

func gameParamsFromPrefs() gameParams {
	return gameParams{
		width:         preferences.IntWithFallback(prefKeys[boardWidthKey], defaultBoardWidth),
		height:        preferences.IntWithFallback(prefKeys[boardHeightKey], defaultBoardHeight),
		infiniteBoard: preferences.BoolWithFallback(prefKeys[infiniteBoardKey], false),
		boardSpeed:    preferences.IntWithFallback(prefKeys[speedKey], 1),
	}
}
