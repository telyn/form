package box

import (
	termbox "github.com/nsf/termbox-go"
)

// TermBox is a box which wraps termbox's cell buffer into a Box
type TermBox struct{}

func (t *TermBox) Size() (int, int) {
	return termbox.Size()
}
func (t *TermBox) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}
