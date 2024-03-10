package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type signal int

const (
	newSpeedSignal signal = iota
	stopSignal
)

type game struct {
	widget.BaseWidget
	*board
	sigChan           chan signal
	patternSelect     *widget.Select
	generationLabel   *widget.Label
	speedRadioButtons *widget.RadioGroup
	playButton        *widget.Button
	clearBoardButton  *widget.Button
	resizeButton      *widget.Button
	speedList         []string
	paused            bool
	speed             int
}

func newGame() *game {
	game := &game{
		sigChan:   make(chan signal),
		board:     newBoard(),
		paused:    true,
		speed:     preferences.IntWithFallback(prefKeys[speedKey], 1),
		speedList: []string{"1x", "5x", "10x", "50x"},
	}
	game.ExtendBaseWidget(game)
	return game
}

func (g *game) buildUI() fyne.CanvasObject {
	g.generationLabel = widget.NewLabel(g.genText())
	g.patternSelect = generatePatterns(g)

	g.playButton = widget.NewButtonWithIcon("PLAY", theme.MediaPlayIcon(), func() {})
	g.playButton.OnTapped = func() {
		if g.paused {
			g.play()
		} else {
			g.pause()
		}
	}

	resetButton := widget.NewButtonWithIcon("RESET", theme.MediaReplayIcon(), func() {
		g.pause()
		g.board.restart()
		g.reset()
	})

	infiniteCheck := widget.NewCheck("Infinite board", func(b bool) {
		g.board.infinite = b
		preferences.SetBool(prefKeys[infiniteBoardKey], b)
	})
	infiniteCheck.SetChecked(preferences.BoolWithFallback(prefKeys[infiniteBoardKey], true))

	g.clearBoardButton = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {})
	g.clearBoardButton.OnTapped = func() {
		dialog.ShowConfirm("Clear board", "This will clear all tiles on board. Continue?", func(confirmed bool) {
			if confirmed {
				g.patternSelect.ClearSelected()
				g.board.initGrid()
				g.reset()
			}
		}, mainWindow)
	}

	g.speedRadioButtons = widget.NewRadioGroup(g.speedList, func(s string) {})
	g.speedRadioButtons.Horizontal = true
	g.speedRadioButtons.Required = true
	g.speedRadioButtons.SetSelected(fmt.Sprintf("%dx", preferences.IntWithFallback(prefKeys[speedKey], 1)))
	g.speedRadioButtons.OnChanged = func(s string) {
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

	heightLabelBind := binding.NewString()
	if err := heightLabelBind.Set(fmt.Sprintf("H: %d", g.board.height)); err != nil {
		dialog.ShowError(err, mainWindow)
	}
	heightLabel := widget.NewLabelWithData(heightLabelBind)

	widthLabelBind := binding.NewString()
	if err := widthLabelBind.Set(fmt.Sprintf("W: %d", g.board.width)); err != nil {
		dialog.ShowError(err, mainWindow)
	}
	widthLabel := widget.NewLabelWithData(widthLabelBind)

	g.resizeButton = widget.NewButton("RESIZE", func() { g.buildResizeDialog().Show() })

	return container.NewBorder(
		container.NewBorder(nil, nil,
			container.NewHBox(
				widget.NewLabel("Pattern:"),
				g.patternSelect,
				widget.NewSeparator(),
				g.playButton,
				resetButton,
				widget.NewSeparator(),
				infiniteCheck,
			),
			container.NewHBox(
				g.clearBoardButton,
			),
		),
		container.NewBorder(nil, nil,
			container.NewHBox(
				widget.NewLabel("Speed:"),
				g.speedRadioButtons,
				widget.NewSeparator(),
				g.generationLabel,
			),
			container.NewHBox(
				g.resizeButton,
				widget.NewSeparator(),
				heightLabel,
				widthLabel,
			),
		),
		nil, nil, g,
	)
}

func (g *game) buildResizeDialog() dialog.Dialog {
	heightEntry := widget.NewEntry()
	heightEntry.Validator = boardSizeValidator
	heightEntry.SetText(fmt.Sprint(g.board.height))
	heightItem := widget.NewFormItem("Height:", heightEntry)
	heightItem.HintText = "Number of board rows"

	widthEntry := widget.NewEntry()
	widthEntry.Validator = boardSizeValidator
	widthEntry.SetText(fmt.Sprint(g.board.width))
	widthItem := widget.NewFormItem("Width:", widthEntry)
	widthItem.HintText = "Number of board columns"

	defaultsButton := widget.NewButton("Reset default", func() {
		heightEntry.SetText(fmt.Sprint(defaultBoardHeight))
		widthEntry.SetText(fmt.Sprint(defaultBoardWidth))
	})
	defaultsItem := widget.NewFormItem("", defaultsButton)

	items := []*widget.FormItem{heightItem, widthItem, defaultsItem}

	return dialog.NewForm("Resize board", "RESIZE", "CANCEL", items, func(confirmed bool) {
		if confirmed {
			h, err := strconv.Atoi(heightEntry.Text)
			if err != nil {
				dialog.ShowError(err, mainWindow)
			}

			w, err := strconv.Atoi(widthEntry.Text)
			if err != nil {
				dialog.ShowError(err, mainWindow)
			}

			preferences.SetInt(prefKeys[boardHeightKey], h)
			preferences.SetInt(prefKeys[boardWidthKey], w)

			g.stop()
			theGame = newGame()
			mainWindow.SetContent(theGame.buildUI())
			mainWindow.SetMainMenu(newMenu(theGame))
			theGame.setKeyPressListener()
			theGame.run()
		}
	}, mainWindow)
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
				switch sig {
				case newSpeedSignal:
					ticker.Reset(time.Second / time.Duration(g.speed))
				case stopSignal:
					g.paused = true
					ticker.Stop()
				}
			}
		}
	}(g.sigChan)
}

func (g *game) stop() {
	g.sigChan <- stopSignal
}

func (g *game) reset() {
	g.board.generation = 0
	g.generationLabel.SetText(g.genText())
	g.Refresh()
}

func (g *game) genText() string {
	return fmt.Sprintf("Generation: %d", g.board.generation)
}

func (g *game) play() {
	g.paused = false
	g.playButton.SetText("PAUSE")
	g.playButton.SetIcon(theme.MediaPauseIcon())
	g.patternSelect.Disable()
	g.resizeButton.Disable()
	g.clearBoardButton.Disable()
}

func (g *game) pause() {
	g.paused = true
	g.playButton.SetText("PLAY")
	g.playButton.SetIcon(theme.MediaPlayIcon())
	g.patternSelect.Enable()
	g.resizeButton.Enable()
	g.clearBoardButton.Enable()
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
	clickedOutsideGrid := event.Position.X < float32(offsetX) ||
		event.Position.Y < float32(offsetY) ||
		event.Position.X > g.Size().Width-float32(offsetX) ||
		event.Position.Y > g.Size().Height-float32(offsetY)

	if clickedOutsideGrid {
		return
	}

	cellX := int(event.Position.X-float32(offsetX)) / g.board.xCellSize
	cellY := int(event.Position.Y-float32(offsetY)) / g.board.yCellSize
	if cellX < g.board.width && cellY < g.board.height {
		g.board.genCurrent[cellY][cellX] = !g.board.genCurrent[cellY][cellX]
	}

	g.reset()
}

func (g *game) TappedSecondary(*fyne.PointEvent) {}

func (g *game) setKeyPressListener() {
	mainWindow.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		index := 0
		for i, v := range g.speedList {
			if v == g.speedRadioButtons.Selected {
				index = i
			}
		}
		if ke.Name == fyne.KeyUp {
			if index+1 < len(g.speedList) {
				index++
			}
		}
		if ke.Name == fyne.KeyDown {
			if index > 0 {
				index--
			}
		}
		g.speedRadioButtons.SetSelected(g.speedList[index])
	})
}
