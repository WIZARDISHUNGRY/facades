// Code generated by a tool. DO NOT EDIT.

// Package dsa provides a mockable wrapper for crypto/dsa.
package dsa

import (
	dsa "crypto/dsa"
	io "io"
	big "math/big"
)

var _ Interface = &Impl{}
var _ = dsa.GenerateKey

type Interface interface {
	GenerateKey(priv *dsa.PrivateKey, rand io.Reader) error
	GenerateParameters(params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) error
	Sign(rand io.Reader, priv *dsa.PrivateKey, hash []byte) (r *big.Int, s *big.Int, err error)
	Verify(pub *dsa.PublicKey, hash []byte, r *big.Int, s *big.Int) bool
}

type Impl struct{}

func (*Impl) GenerateKey(priv *dsa.PrivateKey, rand io.Reader) error {
	return dsa.GenerateKey(priv, rand)
}
func (*Impl) GenerateParameters(params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) error {
	return dsa.GenerateParameters(params, rand, sizes)
}
func (*Impl) Sign(rand io.Reader, priv *dsa.PrivateKey, hash []byte) (r *big.Int, s *big.Int, err error) {
	return dsa.Sign(rand, priv, hash)
}
func (*Impl) Verify(pub *dsa.PublicKey, hash []byte, r *big.Int, s *big.Int) bool {
	return dsa.Verify(pub, hash, r, s)
}