// Code generated by a tool. DO NOT EDIT.

// Package gob provides a mockable wrapper for encoding/gob.
package gob

import (
	gob "encoding/gob"
	io "io"
)

var _ Interface = &Impl{}
var _ = gob.NewDecoder

type Interface interface {
	NewDecoder(r io.Reader) *gob.Decoder
	NewEncoder(w io.Writer) *gob.Encoder
	Register(value any)
	RegisterName(name string, value any)
}

type Impl struct{}

func (*Impl) NewDecoder(r io.Reader) *gob.Decoder {
	return gob.NewDecoder(r)
}
func (*Impl) NewEncoder(w io.Writer) *gob.Encoder {
	return gob.NewEncoder(w)
}
func (*Impl) Register(value any) {
	gob.Register(value)
}
func (*Impl) RegisterName(name string, value any) {
	gob.RegisterName(name, value)
}
