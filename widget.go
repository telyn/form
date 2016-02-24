package form

import (
	"github.com/telyn/form/box"
)

type Widget interface {
	Size() (w int, h int)
	GetCursor() (x int, y int)
	SetCursor(x, y int)
	DrawInto(box box.Box, x int, y int)
	HandleResize(w, h int)
}
