// Code generated by a tool. DO NOT EDIT.

// Package maphash provides a mockable wrapper for hash/maphash.
package maphash

import (
	maphash "hash/maphash"
)

var _ Interface = &Impl{}
var _ = maphash.MakeSeed

type Interface interface {
	MakeSeed() maphash.Seed
}

type Impl struct{}

func (*Impl) MakeSeed() maphash.Seed {
	return maphash.MakeSeed()
}
