package main

import (
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func newMenu(g *game) *fyne.MainMenu {
	return fyne.NewMainMenu(
		newFileMenu(g),
		newPatternMenu(g),
		newBoardMenu(g),
	)
}

func showPauseInfoDialog(window fyne.Window) {
	dialog.ShowInformation("INFO", "You need to pause the game first!", window)
}

func newFileMenu(g *game) *fyne.Menu {
	exportItem := fyne.NewMenuItem("Export board pattern", func() {
		if !g.paused {
			showPauseInfoDialog(mainWindow)
			return
		}
		generateTemplateDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, mainWindow)
				return
			}
			if uc == nil {
				return
			}
			if _, err := uc.Write(generateTemplateBytes(g.board.genCurrent)); err != nil {
				dialog.ShowError(err, mainWindow)
				return
			}
			uc.Close()
		}, mainWindow)
		generateTemplateDialog.Show()
	})

	importItem := fyne.NewMenuItem("Import pattern", func() {
		importTemplateDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, mainWindow)
				return
			}
			if uc == nil {
				return
			}

			// Fyne uc.Read() seems to not be working, so we will ignore it and read file directly.
			// bytes := make([]byte, 0)
			// if _, err := uc.Read(bytes); err != nil {
			// 	dialog.ShowError(err, mainWindow)
			//	return
			// }
			uc.Close()

			bytes, err := os.ReadFile(uc.URI().Path())
			if err != nil {
				dialog.ShowError(err, mainWindow)
				return
			}

			pattern, err := parseImportedPattern(bytes)
			if err != nil {
				dialog.ShowError(err, mainWindow)
				return
			}

			preferences.SetInt(prefKeys[boardHeightKey], len(pattern))
			preferences.SetInt(prefKeys[boardWidthKey], len(pattern[0]))

			g.stop()
			theGame = newGame()
			theGame.board.setStartingPattern(pattern)
			mainWindow.SetContent(theGame.buildUI())
			mainWindow.SetMainMenu(newMenu(theGame))
			theGame.setKeyPressListener()
			theGame.run()
		}, mainWindow)
		importTemplateDialog.Show()
	})

	return fyne.NewMenu("File", exportItem, importItem)
}

func newPatternMenu(g *game) *fyne.Menu {
	patternItems := make([]*fyne.MenuItem, len(patterns))
	for i, name := range patterns {
		patternItems[i] = fyne.NewMenuItem(name, func() {
			g.pause()
			g.board.initGrid()
			midX, midY := g.board.width/2, g.board.height/2
			drawPatternCallback[name](g.board, midX, midY)
			g.board.saveStartPattern()
			g.patternLabel.SetText(strings.ToUpper(name))
			g.reset()
		})
	}

	return fyne.NewMenu("Pattern", patternItems...)
}

func newBoardMenu(g *game) *fyne.Menu {
	clearItem := fyne.NewMenuItem("Clear", func() {
		if !g.paused {
			showPauseInfoDialog(mainWindow)
			return
		}
		g.showClearConfirmDialog()
	})

	resizeItem := fyne.NewMenuItem("Resize\t", func() {
		if !g.paused {
			showPauseInfoDialog(mainWindow)
			return
		}
		g.buildResizeDialog().Show()
	})

	return fyne.NewMenu("Board", clearItem, resizeItem)
}
