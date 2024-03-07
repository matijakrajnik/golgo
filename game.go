package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type signal int

const (
	newSpeedSignal signal = iota
)

type game struct {
	widget.BaseWidget
	*board
	sigChan         chan signal
	patternSelect   *widget.Select
	generationLabel *widget.Label
	paused          bool
	speed           int
}

func newGame() *game {
	game := &game{
		sigChan: make(chan signal),
		board:   newBoard(),
		paused:  true,
		speed:   preferences.IntWithFallback(prefKeys[speedKey], 1),
	}
	game.ExtendBaseWidget(game)
	return game
}

func (g *game) buildUI() fyne.CanvasObject {
	g.generationLabel = widget.NewLabel(g.genText())
	g.patternSelect = generatePatterns(g)

	pauseButton := widget.NewButtonWithIcon("PLAY", theme.MediaPlayIcon(), func() {})
	pauseButton.OnTapped = func() {
		g.paused = !g.paused
		if g.paused {
			pauseButton.SetText("PLAY")
			pauseButton.SetIcon(theme.MediaPlayIcon())
			g.patternSelect.Enable()
		} else {
			pauseButton.SetText("PAUSE")
			pauseButton.SetIcon(theme.MediaPauseIcon())
			g.patternSelect.Disable()
		}
	}

	resetButton := widget.NewButtonWithIcon("RESET", theme.MediaReplayIcon(), func() {
		if !g.paused {
			pauseButton.SetText("PLAY")
			pauseButton.SetIcon(theme.MediaPlayIcon())
		}
		g.paused = true
		g.board.restart()
		g.reset()
	})

	infiniteCheck := widget.NewCheck("Infinite board", func(b bool) {
		g.board.infinite = b
		preferences.SetBool(prefKeys[infiniteBoardKey], b)
	})
	infiniteCheck.SetChecked(preferences.BoolWithFallback(prefKeys[infiniteBoardKey], true))

	speedRadioButtons := widget.NewRadioGroup([]string{"1x", "5x", "10x", "50x"}, func(s string) {})
	speedRadioButtons.Horizontal = true
	speedRadioButtons.Required = true
	speedRadioButtons.SetSelected(fmt.Sprintf("%dx", preferences.IntWithFallback(prefKeys[speedKey], 1)))
	speedRadioButtons.OnChanged = func(s string) {
		re := regexp.MustCompile(`(\d+)x`)
		matches := re.FindStringSubmatch(s)
		if len(matches) < 1 {
			return
		}
		speed, err := strconv.Atoi(matches[1])
		if err != nil {
			return
		}
		g.speed = speed
		preferences.SetInt(prefKeys[speedKey], speed)
		g.sigChan <- newSpeedSignal
	}

	return container.NewBorder(
		container.NewHBox(
			widget.NewLabel("Pattern:"),
			g.patternSelect,
			widget.NewSeparator(),
			widget.NewSeparator(),
			pauseButton,
			resetButton,
			widget.NewSeparator(),
			widget.NewSeparator(),
			infiniteCheck,
		),
		container.NewBorder(nil, nil,
			container.NewHBox(
				widget.NewLabel("Speed:"),
				speedRadioButtons,
			),
			g.generationLabel,
		),
		nil, nil, g,
	)
}

func (g *game) run() {
	go func(sc chan signal) {
		ticker := time.NewTicker(time.Second / time.Duration(g.speed))
		for {
			select {
			case <-ticker.C:
				if g.paused {
					continue
				}
				g.board.nextGen()
				g.generationLabel.SetText(g.genText())
				g.Refresh()
			case sig := <-sc:
				if sig == newSpeedSignal {
					ticker.Reset(time.Second / time.Duration(g.speed))
				}
			}
		}
	}(g.sigChan)
}

func (g *game) reset() {
	g.board.generation = 0
	g.generationLabel.SetText(g.genText())
	g.Refresh()
}

func (g *game) genText() string {
	return fmt.Sprintf("Generation: %d", g.board.generation)
}

func (g *game) CreateRenderer() fyne.WidgetRenderer {
	return newRenderer(g.board)
}

func (g *game) Tapped(event *fyne.PointEvent) {
	if !g.paused {
		return
	}

	if g.patternSelect.SelectedIndex() != -1 {
		tmp := g.board.genCurrent
		g.patternSelect.ClearSelected()
		g.board.genCurrent = tmp
	}

	offsetX, offsetY := g.board.calculateOffset(int(g.Size().Width), int(g.Size().Height))

	cellX := int(event.Position.X-float32(offsetX)) / g.board.xCellSize
	cellY := int(event.Position.Y-float32(offsetY)) / g.board.yCellSize
	if cellX < g.board.width && cellY < g.board.height {
		g.board.genCurrent[cellY][cellX] = !g.board.genCurrent[cellY][cellX]
	}

	g.reset()
}

func (g *game) TappedSecondary(*fyne.PointEvent) {}
