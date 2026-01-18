package cli

import (
	"fmt"
	"io"
)

type Output struct {
	w io.Writer
}

func NewOutput(w io.Writer) *Output {
	return &Output{w: w}
}

func (o *Output) Println(a ...any) {
	_, _ = fmt.Fprintln(o.w, a...)
}

func (o *Output) Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(o.w, format, a...)
}
