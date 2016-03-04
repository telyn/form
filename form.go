package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	"log"
	"time"
)

type Form struct {
	fields       []Field
	currentField int

	escapeSequenceStart time.Time
	escapeSequence      []byte
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

func (f *Form) SelectPreviousField() {
	f.fields[f.currentField].Focus(false)
	f.currentField--
	if f.currentField < 0 {
		f.currentField = len(f.fields) - 1
	}
	f.fields[f.currentField].Focus(true)
	log.Printf("tab recvd. New currentField: %d", f.currentField)
}

func (f *Form) SelectNextField() {
	f.fields[f.currentField].Focus(false)
	f.currentField++
	if f.currentField >= len(f.fields) {
		f.currentField = 0
	}
	f.fields[f.currentField].Focus(true)
	log.Printf("tab recvd. New currentField: %d", f.currentField)
}

func (f *Form) ReceiveKey(key termbox.Key) {
	switch key {
	case termbox.KeyTab:
		f.SelectNextField()
	default:
		f.fields[f.currentField].ReceiveKey(key)
	}
}

// HandleEvent takes a termbox event and
func (f *Form) HandleEvent(ev *termbox.Event) (keepRunning bool) {
	log.Printf("Key: %x Ch: %c", ev.Key, ev.Ch)
	switch ev.Type {
	case termbox.EventKey:
		// we have to deal with escape sequences in this way because Mac terminals are weird
		// see https://github.com/nsf/termbox-go/issues/120
		if time.Since(f.escapeSequenceStart) < 20*time.Millisecond {
			if len(f.escapeSequence) == 1 && f.escapeSequence[0] == byte('[') {
				// basically we're looking for ESC[Z because that's "backwards tab"
				// everything else can be ignored for now.
				if ev.Ch == 'Z' {
					f.SelectPreviousField()
				}
				f.escapeSequence = make([]byte, 0)
				f.escapeSequenceStart = time.Time{}
			} else if len(f.escapeSequence) == 0 && ev.Ch == '[' {
				f.escapeSequence = append(f.escapeSequence, byte(ev.Ch))
			} else {
				f.escapeSequence = make([]byte, 0)
				f.escapeSequenceStart = time.Time{}
				f.HandleEvent(ev)
			}
		} else {
			switch ev.Key {
			case termbox.KeyCtrlC:
				return false
			case termbox.KeyEsc:
				if !f.escapeSequenceStart.IsZero() {
					return false
				}
				f.escapeSequenceStart = time.Now()
			default:
				if ev.Ch == 0x00 {
					f.ReceiveKey(ev.Key)
				} else {
					f.ReceiveRune(ev.Ch)
				}
			}
		}
	case termbox.EventResize:
		f.HandleResize(ev.Width, ev.Height)

	}
	return true
}
