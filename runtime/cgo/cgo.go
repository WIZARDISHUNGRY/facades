// Code generated by a tool. DO NOT EDIT.

// Package cgo provides a mockable wrapper for runtime/cgo.
package cgo

import (
	cgo "runtime/cgo"
)

var _ Interface = &Impl{}
var _ = cgo.NewHandle

type Interface interface {
	NewHandle(v any) cgo.Handle
}

type Impl struct{}

func (*Impl) NewHandle(v any) cgo.Handle {
	return cgo.NewHandle(v)
}
