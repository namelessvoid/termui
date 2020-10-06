package termui

import (
	"image"

	"github.com/gizak/termui/v3/drawille"
)

type Canvas struct {
	Block
	drawille.Canvas
}

func NewCanvas() *Canvas {
	return &Canvas{
		Block:  *NewBlock(),
		Canvas: *drawille.NewCanvas(),
	}
}

func (self *Canvas) SetPoint(p image.Point, color Color) {
	self.Canvas.SetPoint(p, drawille.Color(color))
}

func (self *Canvas) SetLine(p0, p1 image.Point, color Color) {
	self.Canvas.SetLine(p0, p1, drawille.Color(color))
}

func (self *Canvas) SetSprite(position image.Point, sprite *Sprite) {
	for _, point := range sprite.Points {
		if point.X < 0 || point.Y < 0 {
			continue
		}

		self.SetPoint(point.Add(position), ColorWhite)
	}
}

func (self *Canvas) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	for point, cell := range self.Canvas.GetCells() {
		if point.In(self.Rectangle) {
			convertedCell := Cell{
				cell.Rune,
				Style{
					Color(cell.Color),
					ColorClear,
					ModifierClear,
				},
			}
			buf.SetCell(convertedCell, point.Add(self.Inner.Min))
		}
	}
}
