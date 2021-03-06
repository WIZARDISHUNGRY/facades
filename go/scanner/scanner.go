// Code generated by a tool. DO NOT EDIT.

// Package scanner provides a mockable wrapper for go/scanner.
package scanner

import (
	scanner "go/scanner"
	io "io"
)

var _ Interface = &Impl{}
var _ = scanner.PrintError

type Interface interface {
	PrintError(w io.Writer, err error)
}

type Impl struct{}

func (*Impl) PrintError(w io.Writer, err error) {
	scanner.PrintError(w, err)
}
