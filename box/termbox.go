package box

import (
	termbox "github.com/nsf/termbox-go"
)

// TermBox is a box which wraps termbox's cell buffer into a Box
type TermBox struct{}

func (t *TermBox) Fill(cell *termbox.Cell) {
	w, h := termbox.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func (t *TermBox) GetCell(x, y int) *termbox.Cell {
	w, h := termbox.Size()
	idx := y*w + x
	if idx > w*h {
		return nil
	}
	return &termbox.CellBuffer()[idx]
}

func (t *TermBox) Size() (int, int) {
	return termbox.Size()
}
func (t *TermBox) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}
