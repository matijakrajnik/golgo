package main

import (
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func createTestGameParams() gameParams {
	return gameParams{
		width:         defaultBoardWidth,
		height:        defaultBoardHeight,
		infiniteBoard: true,
		boardSpeed:    1,
	}
}

func createTestGame(gp gameParams) *game {
	app := test.NewApp()
	loadPreferences(app)
	mainWindow = app.NewWindow("TEST")
	g := newGame(gp)
	mainWindow.SetContent(g.buildUI())
	mainWindow.SetMainMenu(newMenu(g))
	g.run()
	return g
}

func TestGamePatternLabel(t *testing.T) {
	gp := createTestGameParams()

	t.Run("EmptyByDefault", func(t *testing.T) {
		g := createTestGame(gp)

		assert.Empty(t, g.patternLabel.Text)
	})

	t.Run("SelectedName", func(t *testing.T) {
		g := createTestGame(gp)

		mainWindow.MainMenu().Items[1].Items[0].Action()
		assert.EqualValues(t, "BLINKER", g.patternLabel.Text)
	})

	t.Run("CustomOnChange", func(t *testing.T) {
		g := createTestGame(gp)

		x, y := g.Size().Width/2, g.Size().Height/2
		test.TapAt(g, fyne.Position{X: x, Y: y})
		assert.EqualValues(t, "[CUSTOM PATTERN]", g.patternLabel.Text)
	})
}

func TestGameGenerationLabel(t *testing.T) {
	gp := createTestGameParams()

	t.Run("ZeroByDefault", func(t *testing.T) {
		g := createTestGame(gp)

		assert.EqualValues(t, "Generation: 0", g.generationLabel.Text)
	})

	t.Run("IncreasesOnTick", func(t *testing.T) {
		g := createTestGame(gp)
		mainWindow.MainMenu().Items[1].Items[0].Action()
		test.Tap(g.playButton)
		time.Sleep(1050 * time.Millisecond)
		test.Tap(g.playButton)

		assert.EqualValues(t, "Generation: 1", g.generationLabel.Text)
	})
}

func TestPlayPauseButton(t *testing.T) {
	gp := createTestGameParams()
	g := createTestGame(gp)

	test.Tap(g.playButton)
	time.Sleep(1050 * time.Millisecond)
	test.Tap(g.playButton)

	assert.EqualValues(t, "Generation: 1", g.generationLabel.Text)
	time.Sleep(3 * time.Second)
	assert.EqualValues(t, "Generation: 1", g.generationLabel.Text)
	test.Tap(g.playButton)
	time.Sleep(1050 * time.Millisecond)
	assert.EqualValues(t, "Generation: 2", g.generationLabel.Text)
}

func TestSpeedRadioButton(t *testing.T) {
	gp := createTestGameParams()
	g := createTestGame(gp)

	g.speedRadioButtons.SetSelected("10x")
	test.Tap(g.playButton)
	time.Sleep(1050 * time.Millisecond)
	test.Tap(g.playButton)

	assert.EqualValues(t, "Generation: 10", g.generationLabel.Text)
}
