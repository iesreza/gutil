package str

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type str struct {
	s string
}

// S create string instance from any object
func S(v interface{}) *str {
	obj := str{}
	obj.s = fmt.Sprintf("%v", v)
	return &obj
}

func New(v interface{}) *str {
	obj := str{}
	obj.s = fmt.Sprintf("%v", v)
	return &obj
}

// String return string value
func (o *str) String() string {
	return o.s
}

// Quote quote string
func (o *str) Quote() *str {
	o.s = strconv.Quote(o.s)
	return o
}

// TrimSpace
func (o *str) TrimSpace() *str {
	o.s = strings.TrimSpace(o.s)
	return o
}

// Ltrim
func (o *str) Ltrim(cutset string) *str {
	o.s = strings.TrimLeft(o.s, cutset)
	return o
}

// Rtrim
func (o *str) Rtrim(cutset string) *str {
	o.s = strings.TrimRight(o.s, cutset)
	return o
}

// Trim
func (o *str) Trim(cutset string) *str {
	o.s = strings.Trim(o.s, cutset)
	return o
}

// Replace
func (o *str) Replace(what, with string) *str {
	o.s = strings.Replace(o.s, what, with, 1)
	return o
}

// ReplaceAll
func (o *str) ReplaceAll(what, with string) *str {
	o.s = strings.Replace(o.s, what, with, -1)
	return o
}

// IsEmpty
func (o *str) IsEmpty() bool {
	return len(o.s) == 0
}

// IsNotEmpty returns true if the string is not empty
func (o *str) IsNotEmpty() bool {
	return !o.IsEmpty()
}

// IsBlank returns true if the string is blank (all whitespace)
func (o *str) IsBlank() bool {
	return len(strings.TrimSpace(o.s)) == 0
}

// IsNotBlank returns true if the string is not blank
func (o *str) IsNotBlank() bool {
	return !o.IsBlank()
}

// Left justifies the text to the left
func (o *str) Left(size int) *str {
	spaces := size - len(o.s)
	if spaces <= 0 {
		return o
	}

	var buffer bytes.Buffer
	buffer.WriteString(o.s)

	for i := 0; i < spaces; i++ {
		buffer.WriteString(" ")
	}
	o.s = buffer.String()
	return o
}

// Right justifies the text to the right
func (o *str) Right(size int) *str {
	spaces := size - len(o.s)
	if spaces <= 0 {
		return o
	}

	var buffer bytes.Buffer
	for i := 0; i < spaces; i++ {
		buffer.WriteString(" ")
	}

	buffer.WriteString(o.s)
	o.s = buffer.String()
	return o
}
