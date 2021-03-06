// Code generated by a tool. DO NOT EDIT.

// Package dwarf provides a mockable wrapper for debug/dwarf.
package dwarf

import (
	dwarf "debug/dwarf"
)

var _ Interface = &Impl{}
var _ = dwarf.New

type Interface interface {
	New(abbrev []byte, aranges []byte, frame []byte, info []byte, line []byte, pubnames []byte, ranges []byte, str []byte) (*dwarf.Data, error)
}

type Impl struct{}

func (*Impl) New(abbrev []byte, aranges []byte, frame []byte, info []byte, line []byte, pubnames []byte, ranges []byte, str []byte) (*dwarf.Data, error) {
	return dwarf.New(abbrev, aranges, frame, info, line, pubnames, ranges, str)
}
