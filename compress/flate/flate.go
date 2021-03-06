// Code generated by a tool. DO NOT EDIT.

// Package flate provides a mockable wrapper for compress/flate.
package flate

import (
	flate "compress/flate"
	io "io"
)

var _ Interface = &Impl{}
var _ = flate.NewReader

type Interface interface {
	NewReader(r io.Reader) io.ReadCloser
	NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
	NewWriter(w io.Writer, level int) (*flate.Writer, error)
	NewWriterDict(w io.Writer, level int, dict []byte) (*flate.Writer, error)
}

type Impl struct{}

func (*Impl) NewReader(r io.Reader) io.ReadCloser {
	return flate.NewReader(r)
}
func (*Impl) NewReaderDict(r io.Reader, dict []byte) io.ReadCloser {
	return flate.NewReaderDict(r, dict)
}
func (*Impl) NewWriter(w io.Writer, level int) (*flate.Writer, error) {
	return flate.NewWriter(w, level)
}
func (*Impl) NewWriterDict(w io.Writer, level int, dict []byte) (*flate.Writer, error) {
	return flate.NewWriterDict(w, level, dict)
}
