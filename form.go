package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	"log"
)

type Form struct {
	fields       []Field
	currentField int
}

func NewForm(fields []Field) (form *Form) {
	form = new(Form)
	for _, f := range fields {
		form.AddField(f)
	}
	return form
}

func (f *Form) AddField(field Field) {
	bLen := len(f.fields)
	if bLen == 0 {
		field.Focus(true)
	}
	f.fields = append(f.fields, field)
	log.Printf("Adding a field. Before len: %d after len: %d", bLen, len(f.fields))
}

func (f *Form) HandleResize(w, h int) {
	for _, field := range f.fields {
		field.HandleResize(w, h)
	}
}

func (f *Form) DrawInto(box box.Box, offsetX, offsetY int) {
	log.Printf("Now that we're drawing, %d fields", len(f.fields))
	if len(f.fields) == 0 {
		return
	}

	boxW, boxH := box.Size()

	currentY := offsetY

	for i, field := range f.fields {
		field.HandleResize(boxW-offsetX, boxH-offsetY)
		_, fieldH := field.Size()
		log.Printf("field %d: box: %dx%d offset: (%d,%d) fieldH: %d)", i, boxW, boxH, offsetX, currentY, fieldH)
		if currentY+fieldH > boxH {
			return
		}
		field.DrawInto(box, offsetX, currentY)
		currentY += fieldH + 1
	}
}

func (f *Form) ReceiveRune(ch rune) {
	f.fields[f.currentField].ReceiveRune(ch)
}
func (f *Form) ReceiveKey(key termbox.Key) {
	f.fields[f.currentField].ReceiveKey(key)
}
