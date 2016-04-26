package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	"strings"
)

type LabelledField struct {
	innerField   Field
	label        string
	errorMessage string
	outerWidth   int // used for Size()
}

func Label(f Field, label string) Field {
	return &LabelledField{
		innerField: f,
		label:      label,
	}
}

func (f *LabelledField) DrawInto(box box.Box, offsetX, offsetY int) {
	boxW, _ /*boxH*/ := box.Size()
	sizeX, _ /*sizeY*/ := f.Size()
	innerWidth, _ := f.innerField.Size()

	// | label field |
	if offsetX+sizeX > boxW {
		// it aint gonna fit.
		return
	}

	// -2 for spaces either side of the label, -3 for space after the innerField
	labelWidth := (boxW - offsetX) - innerWidth - 5
	DrawString(f.label, box, offsetX+1, offsetY, labelWidth)
	f.innerField.DrawInto(box, offsetX+labelWidth+4, offsetY)

	return
}

func (f *LabelledField) Focus(hasFocus bool) {
	f.innerField.Focus(hasFocus)
}

func (f *LabelledField) GetCursor() (x, y int) {
	return -1, 0
}

func (f *LabelledField) HandleResize(w, h int) {
	f.outerWidth = w
}

func (f *LabelledField) ReceiveKey(key termbox.Key) {
	f.innerField.ReceiveKey(key)
}
func (f *LabelledField) ReceiveRune(ch rune) {
	f.innerField.ReceiveRune(ch)
}

func (f *LabelledField) Size() (w, h int) {
	fieldW, fieldH := f.innerField.Size()
	labelWidth := f.outerWidth - fieldW - 5

	str := FlowString(f.label, labelWidth)
	lines := strings.Count(str, "\n") + 1
	if lines < fieldH {
		lines = fieldH
	}

	return f.outerWidth, lines
}

func (f *LabelledField) Validate() (string, bool) {
	prob, ok := f.innerField.Validate()
	upTo := strings.IndexAny(f.label, "\r(")
	if upTo == -1 {
		prob = f.label + ": " + prob
	} else {
		prob = strings.TrimSpace(f.label[0:upTo]) + ": " + prob
	}
	return prob, ok

}

func (f *LabelledField) Value() string {
	return f.innerField.Value()
}

func (f *LabelledField) SetCursor(x, y int) {
	return
}

func GetCursor() (x int, y int) {
	return
}
