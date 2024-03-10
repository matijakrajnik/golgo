package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func newMenu(g *game) *fyne.MainMenu {
	exportMenuItem := fyne.NewMenuItem("Export board pattern", func() {
		g.pause()
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

	importMenuItem := fyne.NewMenuItem("Import pattern", func() {
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

	fileMenu := fyne.NewMenu("File", exportMenuItem, importMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	return menu
}
