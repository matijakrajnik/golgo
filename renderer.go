package main

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

const minCellSize = 8

type renderer struct {
	*board
	raster        *canvas.Raster
	objects       []fyne.CanvasObject
	colorAlive    color.Color
	colorDead     color.Color
	colorOverflow color.Color
}

func newRenderer(board *board) *renderer {
	renderer := &renderer{
		board:         board,
		objects:       make([]fyne.CanvasObject, 0),
		colorAlive:    color.White,
		colorDead:     color.Black,
		colorOverflow: theme.BackgroundColor(),
	}
	renderer.raster = canvas.NewRaster(renderer.drawImage)
	renderer.objects = append(renderer.objects, renderer.raster)

	return renderer
}

func (r *renderer) drawImage(w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	offsetX, offsetY := r.board.calculateOffset(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, r.colorOverflow)
			if x < offsetX || x > +w-offsetX || y < offsetY || y > +h-offsetY {
				continue
			}
			cellX := (x - offsetX) / r.board.xCellSize
			cellY := (y - offsetY) / r.board.yCellSize
			if cellX < r.board.width && cellY < r.board.height && r.board.genCurrent[cellY][cellX] {
				img.Set(x, y, r.colorAlive)
			} else {
				img.Set(x, y, r.colorDead)
			}
		}
	}

	return img
}

func (r *renderer) Layout(size fyne.Size) {
	r.board.calculateCellSize(int(size.Width), int(size.Height))
	r.raster.Resize(size)
}

func (r *renderer) MinSize() fyne.Size {
	return fyne.NewSize(float32(r.board.width*minCellSize), float32(r.board.height*minCellSize))
}

func (r *renderer) Refresh() {
	r.raster.Refresh()
}

func (r *renderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *renderer) Destroy() {}
