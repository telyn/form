package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	"log"
)

type LabelledField struct {
	innerField   Field
	label        string
	errorMessage string
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

	labelWidth := boxW - innerWidth - 3
	log.Printf("labelWidth: %d", labelWidth)
	DrawString(f.label, box, offsetX+1, offsetY, labelWidth)
	return
}

func (f *LabelledField) Focus(hasFocus bool) {
	f.innerField.Focus(hasFocus)
}

func (f *LabelledField) GetCursor() (x, y int) {
	return -1, 0
}

func (f *LabelledField) ReceiveKey(key termbox.Key) {
	f.innerField.ReceiveKey(key)
}
func (f *LabelledField) ReceiveRune(ch rune) {
	f.innerField.ReceiveRune(ch)
}

func (f *LabelledField) Size() (w, h int) {
	fieldW, fieldH := f.innerField.Size()
	return fieldW + 7, fieldH
}

func (f *LabelledField) Validate() bool {
	return f.innerField.Validate()
}

func (f *LabelledField) SetCursor(x, y int) {
	return
}

func GetCursor() (x int, y int) {
	return
}
