package asm

import "fmt"

type Output struct {
	lines []string
}

func newOutput() *Output {
	return &Output{
		lines: make([]string, 0),
	}
}

func (o *Output) Directive(name string, content string) {
	o.lines = append(o.lines, fmt.Sprintf("    .%s  %s", name, content))
}

func (o *Output) Label(label string) {
	o.lines = append(o.lines, fmt.Sprintf("%s:", label))
}
