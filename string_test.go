package form

import (
	"fmt"
	"github.com/cheekybits/is"
	"os"
	"testing"
)

func TestFlowString(t *testing.T) {
	is := is.New(t)
	testInputs := []string{
		`hey hey you yeah i don't like your girlfriend no way no way i think you need a new one`,
	}
	testOutputs := []string{
		"hey hey you yeah i don't like\r\nyour girlfriend no way no way\r\ni think you need a new one",
	}
	for i, in := range testInputs {
		fmt.Fprintln(os.Stderr, []byte(testOutputs[i]))
		fmt.Fprintln(os.Stderr, []byte(FlowString(in, 30)))
		is.Equal(testOutputs[i], FlowString(in, 30))
	}
}
