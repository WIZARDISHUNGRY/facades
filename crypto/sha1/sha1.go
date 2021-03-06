// Code generated by a tool. DO NOT EDIT.

// Package sha1 provides a mockable wrapper for crypto/sha1.
package sha1

import (
	sha1 "crypto/sha1"
	hash "hash"
)

var _ Interface = &Impl{}
var _ = sha1.New

type Interface interface {
	New() hash.Hash
	Sum(data []byte) [20]byte
}

type Impl struct{}

func (*Impl) New() hash.Hash {
	return sha1.New()
}
func (*Impl) Sum(data []byte) [20]byte {
	return sha1.Sum(data)
}
