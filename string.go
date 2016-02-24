package form

import (
	"github.com/telyn/form/box"
	"log"
	"strings"
)

func FlowString(str string, width int) string {
	words := strings.Split(str, " ")
	lines := make([]string, 0)
	line := ""
	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(line)+1+len(word) > width {
			lines = append(lines, line)
			line = word
		} else {
			if line == "" {
				line = word
			} else {
				line += " " + word
			}
		}
	}
	if line != "" {
		return strings.Join(lines, "\r\n") + "\r\n" + line
	} else {
		return strings.Join(lines, "\r\n")
	}

}

func DrawString(str string, box box.Box, offsetX, offsetY, width int) {
	curX := offsetX
	curY := offsetY
	str = FlowString(str, width)
	for _, ch := range str {
		switch ch {
		case '\r':
			curX = offsetX
		case '\n':
			curY = offsetY + 1
		default:
			cell := box.GetCell(curX, curY)
			log.Printf("box.SetCell(%d,%d,%c,fg,bg)", curX, curY, ch)
			box.SetCell(curX, curY, ch, cell.Fg, cell.Bg)
		}
		curX++
	}

}
