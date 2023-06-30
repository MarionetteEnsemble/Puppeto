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

func (p Position) to_string() string {
	return fmt.Sprintf("%s(%s, %s)", p.File, strconv.FormatInt(p.Line, 10), strconv.FormatInt(p.Column, 10))
}

type Stack = []Position

func Stackto_string(stack Stack) string {
	ss := []string{}

	for _, pos := range stack {
		ss = append(ss, "\tat "+pos.to_string())
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
	o["to_string"] = func(stack Stack, args []any) (any, any) {
		return fmt.Sprintf("%s: %s\n%s", o["name"], message, Stackto_string(stack)), nil
	}

	return o
}
