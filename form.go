package form

import (
	"errors"
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	//"log"
	"time"
)

type Form struct {
	fields          []Field
	currentField    int
	currentTopField int

	escapeSequenceStart time.Time
	escapeSequence      []byte
	running             bool
}

func NewForm(fields []Field) (form *Form) {
	form = new(Form)
	for _, f := range fields {
		form.AddField(f)
	}
	form.currentField = -1
	form.SelectNextField()
	return form
}

func (f *Form) AddField(field Field) {
	bLen := len(f.fields)
	if bLen == 0 {
		field.Focus(true)
	}
	f.fields = append(f.fields, field)
}

func (f *Form) HandleResize(w, h int) {
	for _, field := range f.fields {
		field.HandleResize(w, h)
	}
}

func (f *Form) DrawInto(box box.Box, offsetX, offsetY int) {
	if len(f.fields) == 0 {
		return
	}

	boxW, boxH := box.Size()

	currentY := offsetY
	for _, field := range f.fields {
		field.HandleResize(boxW-offsetX, boxH-offsetY)
	}

	// it would be nice to draw into an infinitely large box and then only copy the relevant portion.. this architecture doesn't really allow for that though
	f.ensureCurrentFieldOnScreen(boxH)

	for _, field := range f.fields[f.currentTopField:] {
		_, fieldH := field.Size()
		//log.Printf("field %d: box: %dx%d offset: (%d,%d) fieldH: %d)", i, boxW, boxH, offsetX, currentY, fieldH)
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

func (f *Form) Run() error {
	draw := func() bool {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		f.DrawInto(&box.TermBox{}, 0, 0)

		termbox.Flush()
		return true
	}
	err := termbox.Init()
	defer termbox.Close()
	if err != nil {
		return err
	}

	if draw() {
		f.running = true
		for f.running {
			ev := termbox.PollEvent()
			if !f.HandleEvent(&ev) {
				f.running = false
			}
			draw()
		}
	} else {
		return errors.New("fail sadness")
	}
	return nil
}

func (f *Form) Stop() {
	f.running = false
}

func (f *Form) ensureCurrentFieldOnScreen(boxH int) {
	//log.Printf("OLD currentField: %d currentTopField: %d boxH: %d", f.currentField, f.currentTopField, boxH)
	// scroll down ONLY AS FAR AS NECESSARY
	// but we need to know the box size in order to do so.

	// start at current field and work way back up to find top field
	height := 0
	// this will do weird crap when the window is too small to fit the currentField, but never mind... i guess

	// loop backwards from target field until we find a suitable starting point
	top := 0
	for top = f.currentField; top > -1; top-- {
		_, h := f.fields[top].Size()
		height += h + 1
		//log.Printf("iter top: %d height: %d h: %d", top, height, h)
		if height >= boxH {
			f.currentTopField = top + 1
			//log.Printf("NEW currentField: %d currentTopField: %d", f.currentField, f.currentTopField)
			return
		}
	}

	f.currentTopField = 0

	//log.Printf("NEW currentField: %d currentTopField: %d", f.currentField, f.currentTopField)
}

func (f *Form) SelectPreviousField() {
	f.fields[f.currentField].Focus(false)
	f.currentField--
	if f.currentField < 0 {
		f.currentField = len(f.fields) - 1
	}
	f.fields[f.currentField].Focus(true)
	if _, ok := f.fields[f.currentField].(*LabelField); ok {
		f.SelectPreviousField()
	}
	//log.Printf("tab recvd. New currentField: %d", f.currentField)
}

func (f *Form) SelectNextField() {
	if f.currentField >= 0 {
		f.fields[f.currentField].Focus(false)
	}
	f.currentField++
	if f.currentField >= len(f.fields) {
		f.currentField = 0
	}
	f.fields[f.currentField].Focus(true)
	if _, ok := f.fields[f.currentField].(*LabelField); ok {
		f.SelectNextField()
	}
	//log.Printf("tab recvd. New currentField: %d", f.currentField)
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
	//log.Printf("Key: %x Ch: %c", ev.Key, ev.Ch)
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

func (f *Form) Validate() ([]string, bool) {
	problems := make([]string, 0)
	for _, field := range f.fields {
		if p, ok := field.Validate(); !ok {
			problems = append(problems, p)
		}
	}
	return problems, len(problems) == 0
}
