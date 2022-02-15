// Code generated by a tool. DO NOT EDIT.

// Package des provides a mockable wrapper for crypto/des.
package des

import (
	cipher "crypto/cipher"
	des "crypto/des"
)

var _ Interface = &Impl{}
var _ = des.NewCipher

type Interface interface {
	NewCipher(key []byte) (cipher.Block, error)
	NewTripleDESCipher(key []byte) (cipher.Block, error)
}

type Impl struct{}

func (*Impl) NewCipher(key []byte) (cipher.Block, error) {
	return des.NewCipher(key)
}
func (*Impl) NewTripleDESCipher(key []byte) (cipher.Block, error) {
	return des.NewTripleDESCipher(key)
}