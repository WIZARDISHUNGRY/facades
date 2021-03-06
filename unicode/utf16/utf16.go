// Code generated by a tool. DO NOT EDIT.

// Package utf16 provides a mockable wrapper for unicode/utf16.
package utf16

import (
	utf16 "unicode/utf16"
)

var _ Interface = &Impl{}
var _ = utf16.Decode

type Interface interface {
	Decode(s []uint16) []rune
	DecodeRune(r1 rune, r2 rune) rune
	Encode(s []rune) []uint16
	EncodeRune(r rune) (r1 rune, r2 rune)
	IsSurrogate(r rune) bool
}

type Impl struct{}

func (*Impl) Decode(s []uint16) []rune {
	return utf16.Decode(s)
}
func (*Impl) DecodeRune(r1 rune, r2 rune) rune {
	return utf16.DecodeRune(r1, r2)
}
func (*Impl) Encode(s []rune) []uint16 {
	return utf16.Encode(s)
}
func (*Impl) EncodeRune(r rune) (r1 rune, r2 rune) {
	return utf16.EncodeRune(r)
}
func (*Impl) IsSurrogate(r rune) bool {
	return utf16.IsSurrogate(r)
}
