package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	"strings"
)

type LabelField struct {
	label      string
	outerWidth int // used for Size()
}

func NewLabelField(label string) (lf *LabelField) {
	lf = new(LabelField)
	lf.label = label
	return
}

func (f *LabelField) DrawInto(box box.Box, offsetX, offsetY int) {
	boxW, _ /*boxH*/ := box.Size()
	sizeX, _ /*sizeY*/ := f.Size()

	// | label field |
	if offsetX+sizeX > boxW {
		// it aint gonna fit.
		return
	}

	// -2 for spaces either side of the label, -3 for space after the innerField
	DrawString(f.label, box, offsetX+1, offsetY, boxW-offsetX-2)

	return
}

func (f *LabelField) Focus(hasFocus bool) {
	return
}

func (f *LabelField) GetCursor() (x, y int) {
	return -1, 0
}

func (f *LabelField) HandleResize(w, h int) {
	f.outerWidth = w
}

func (f *LabelField) ReceiveKey(key termbox.Key) {
	return
}
func (f *LabelField) ReceiveRune(ch rune) {
	return
}

func (f *LabelField) Size() (w, h int) {
	labelWidth := f.outerWidth - 2

	str := FlowString(f.label, labelWidth)
	lines := strings.Count(str, "\n") + 1

	return f.outerWidth, lines
}

func (f *LabelField) Validate() (string, bool) {
	return "", true
}

func (f *LabelField) Value() string {
	return ""
}

func (f *LabelField) SetCursor(x, y int) {
	return
}
