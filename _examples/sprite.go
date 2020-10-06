package main

import (
	"image"
	"log"

	ui "github.com/gizak/termui/v3"
)

func main() {
	err := ui.Init()
	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	canvas := ui.NewCanvas()
	canvas.SetRect(0, 0, 50, 50)

	sprite := ui.NewSprite()
	sprite.Points = []image.Point{
		image.Pt(0, 0),
		image.Pt(1, 1),
		image.Pt(2, 2),
		image.Pt(3, 3),
		image.Pt(4, 4),
		image.Pt(4, 0),
		image.Pt(3, 1),
		image.Pt(1, 3),
		image.Pt(0, 4),
	}

	canvas.SetSprite(image.Pt(2, 2), sprite)
	ui.Render(canvas)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
