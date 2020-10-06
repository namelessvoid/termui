package widgets

import (
	"image"

	ui "github.com/gizak/termui/v3"
)

type RuneCanvas struct {
	ui.Block
	cellMap map[image.Point]ui.Cell
}

func NewRuneCanvas() *RuneCanvas {
	r := &RuneCanvas{
		Block: *ui.NewBlock(),
	}
	r.Clear()
	return r
}

func (r *RuneCanvas) Clear() {
	r.cellMap = make(map[image.Point]ui.Cell)
}

func (r *RuneCanvas) SetCell(point image.Point, rune rune, colors ...interface{}) {
	r.cellMap[point] = ui.NewCell(rune, colors...)
}

func (r *RuneCanvas) GetCell(point image.Point) ui.Cell {
	return r.cellMap[point]
}

func (r *RuneCanvas) GetCells() map[image.Point]ui.Cell {
	return r.cellMap
}

func (r *RuneCanvas) SetCellBgColor(point image.Point, bgColor ui.Color) {
	cell := r.cellMap[point]
	cell.Style.Bg = bgColor
	r.cellMap[point] = cell
}

func (r *RuneCanvas) SetSprite(position image.Point, sprite *ui.Sprite, rune rune, style ui.Style) {
	for _, point := range sprite.Points {
		r.SetCell(point.Add(position), rune, style)
	}
}

func (r *RuneCanvas) Draw(buffer *ui.Buffer) {
	r.Block.Draw(buffer)
	for point, cell := range r.cellMap {
		if point.X < 0 || point.Y < 0 || point.X > r.Inner.Dx()-1 || point.Y > r.Inner.Dy()-1 {
			continue
		}

		buffer.SetCell(cell, point.Add(r.Inner.Min))
	}
}
