// Code generated by a tool. DO NOT EDIT.

// Package pem provides a mockable wrapper for encoding/pem.
package pem

import (
	pem "encoding/pem"
	io "io"
)

var _ Interface = &Impl{}
var _ = pem.Decode

type Interface interface {
	Decode(data []byte) (p *pem.Block, rest []byte)
	Encode(out io.Writer, b *pem.Block) error
	EncodeToMemory(b *pem.Block) []byte
}

type Impl struct{}

func (*Impl) Decode(data []byte) (p *pem.Block, rest []byte) {
	return pem.Decode(data)
}
func (*Impl) Encode(out io.Writer, b *pem.Block) error {
	return pem.Encode(out, b)
}
func (*Impl) EncodeToMemory(b *pem.Block) []byte {
	return pem.EncodeToMemory(b)
}
