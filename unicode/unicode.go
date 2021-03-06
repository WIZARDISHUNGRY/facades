// Code generated by a tool. DO NOT EDIT.

// Package unicode provides a mockable wrapper for unicode.
package unicode

import (
	unicode "unicode"
)

var _ Interface = &Impl{}
var _ = unicode.In

type Interface interface {
	In(r rune, ranges ...*unicode.RangeTable) bool
	Is(rangeTab *unicode.RangeTable, r rune) bool
	IsControl(r rune) bool
	IsDigit(r rune) bool
	IsGraphic(r rune) bool
	IsLetter(r rune) bool
	IsLower(r rune) bool
	IsMark(r rune) bool
	IsNumber(r rune) bool
	IsOneOf(ranges []*unicode.RangeTable, r rune) bool
	IsPrint(r rune) bool
	IsPunct(r rune) bool
	IsSpace(r rune) bool
	IsSymbol(r rune) bool
	IsTitle(r rune) bool
	IsUpper(r rune) bool
	SimpleFold(r rune) rune
	To(_case int, r rune) rune
	ToLower(r rune) rune
	ToTitle(r rune) rune
	ToUpper(r rune) rune
}

type Impl struct{}

func (*Impl) In(r rune, ranges ...*unicode.RangeTable) bool {
	return unicode.In(r, ranges...)
}
func (*Impl) Is(rangeTab *unicode.RangeTable, r rune) bool {
	return unicode.Is(rangeTab, r)
}
func (*Impl) IsControl(r rune) bool {
	return unicode.IsControl(r)
}
func (*Impl) IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}
func (*Impl) IsGraphic(r rune) bool {
	return unicode.IsGraphic(r)
}
func (*Impl) IsLetter(r rune) bool {
	return unicode.IsLetter(r)
}
func (*Impl) IsLower(r rune) bool {
	return unicode.IsLower(r)
}
func (*Impl) IsMark(r rune) bool {
	return unicode.IsMark(r)
}
func (*Impl) IsNumber(r rune) bool {
	return unicode.IsNumber(r)
}
func (*Impl) IsOneOf(ranges []*unicode.RangeTable, r rune) bool {
	return unicode.IsOneOf(ranges, r)
}
func (*Impl) IsPrint(r rune) bool {
	return unicode.IsPrint(r)
}
func (*Impl) IsPunct(r rune) bool {
	return unicode.IsPunct(r)
}
func (*Impl) IsSpace(r rune) bool {
	return unicode.IsSpace(r)
}
func (*Impl) IsSymbol(r rune) bool {
	return unicode.IsSymbol(r)
}
func (*Impl) IsTitle(r rune) bool {
	return unicode.IsTitle(r)
}
func (*Impl) IsUpper(r rune) bool {
	return unicode.IsUpper(r)
}
func (*Impl) SimpleFold(r rune) rune {
	return unicode.SimpleFold(r)
}
func (*Impl) To(_case int, r rune) rune {
	return unicode.To(_case, r)
}
func (*Impl) ToLower(r rune) rune {
	return unicode.ToLower(r)
}
func (*Impl) ToTitle(r rune) rune {
	return unicode.ToTitle(r)
}
func (*Impl) ToUpper(r rune) rune {
	return unicode.ToUpper(r)
}
