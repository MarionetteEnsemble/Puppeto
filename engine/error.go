package engine

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	Line   int64
	Column int64
	File   string
}

func NewPosition(line int64, column int64, file string) Position {
	return Position{
		Line:   line,
		Column: column,
		File:   file,
	}
}

func (p Position) ToString() string {
	return fmt.Sprintf("%s(%s, %s)", p.File, strconv.FormatInt(p.Line, 10), strconv.FormatInt(p.Column, 10))
}

type Stack = []Position

func StackToString(stack Stack) string {
	ss := []string{}

	for _, pos := range stack {
		ss = append(ss, "\tat "+pos.ToString())
	}

	return strings.Join(ss, "\n")
}

const (
	TypeError   = "TypeError"
	SyntaxError = "SyntaxError"
)

func NewPError(name string, message string, stack Stack) PError {
	o := PObject{}

	o["name"] = name
	o["message"] = message
	o["stack"] = stack
	o["toString"] = func() string {
		return fmt.Sprintf("%s: %s\n%s", o["name"], message, StackToString(stack))
	}

	return o
}
