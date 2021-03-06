// Code generated by a tool. DO NOT EDIT.

// Package x509 provides a mockable wrapper for crypto/x509.
package x509

import (
	crypto "crypto"
	ecdsa "crypto/ecdsa"
	rsa "crypto/rsa"
	x509 "crypto/x509"
	pkix "crypto/x509/pkix"
	pem "encoding/pem"
	io "io"
)

var _ Interface = &Impl{}
var _ = x509.CreateCertificate

type Interface interface {
	CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error)
	CreateCertificateRequest(rand io.Reader, template *x509.CertificateRequest, priv any) (csr []byte, err error)
	CreateRevocationList(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) ([]byte, error)
	DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error)
	EncryptPEMBlock(rand io.Reader, blockType string, data []byte, password []byte, alg x509.PEMCipher) (*pem.Block, error)
	IsEncryptedPEMBlock(b *pem.Block) bool
	MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)
	MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
	MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte
	MarshalPKCS8PrivateKey(key any) ([]byte, error)
	MarshalPKIXPublicKey(pub any) ([]byte, error)
	NewCertPool() *x509.CertPool
	ParseCRL(crlBytes []byte) (*pkix.CertificateList, error)
	ParseCertificate(der []byte) (*x509.Certificate, error)
	ParseCertificateRequest(asn1Data []byte) (*x509.CertificateRequest, error)
	ParseCertificates(der []byte) ([]*x509.Certificate, error)
	ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error)
	ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error)
	ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)
	ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)
	ParsePKCS8PrivateKey(der []byte) (key any, err error)
	ParsePKIXPublicKey(derBytes []byte) (pub any, err error)
	SystemCertPool() (*x509.CertPool, error)
}

type Impl struct{}

func (*Impl) CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error) {
	return x509.CreateCertificate(rand, template, parent, pub, priv)
}
func (*Impl) CreateCertificateRequest(rand io.Reader, template *x509.CertificateRequest, priv any) (csr []byte, err error) {
	return x509.CreateCertificateRequest(rand, template, priv)
}
func (*Impl) CreateRevocationList(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) ([]byte, error) {
	return x509.CreateRevocationList(rand, template, issuer, priv)
}
func (*Impl) DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error) {
	return x509.DecryptPEMBlock(b, password)
}
func (*Impl) EncryptPEMBlock(rand io.Reader, blockType string, data []byte, password []byte, alg x509.PEMCipher) (*pem.Block, error) {
	return x509.EncryptPEMBlock(rand, blockType, data, password, alg)
}
func (*Impl) IsEncryptedPEMBlock(b *pem.Block) bool {
	return x509.IsEncryptedPEMBlock(b)
}
func (*Impl) MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	return x509.MarshalECPrivateKey(key)
}
func (*Impl) MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}
func (*Impl) MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte {
	return x509.MarshalPKCS1PublicKey(key)
}
func (*Impl) MarshalPKCS8PrivateKey(key any) ([]byte, error) {
	return x509.MarshalPKCS8PrivateKey(key)
}
func (*Impl) MarshalPKIXPublicKey(pub any) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(pub)
}
func (*Impl) NewCertPool() *x509.CertPool {
	return x509.NewCertPool()
}
func (*Impl) ParseCRL(crlBytes []byte) (*pkix.CertificateList, error) {
	return x509.ParseCRL(crlBytes)
}
func (*Impl) ParseCertificate(der []byte) (*x509.Certificate, error) {
	return x509.ParseCertificate(der)
}
func (*Impl) ParseCertificateRequest(asn1Data []byte) (*x509.CertificateRequest, error) {
	return x509.ParseCertificateRequest(asn1Data)
}
func (*Impl) ParseCertificates(der []byte) ([]*x509.Certificate, error) {
	return x509.ParseCertificates(der)
}
func (*Impl) ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error) {
	return x509.ParseDERCRL(derBytes)
}
func (*Impl) ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error) {
	return x509.ParseECPrivateKey(der)
}
func (*Impl) ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(der)
}
func (*Impl) ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(der)
}
func (*Impl) ParsePKCS8PrivateKey(der []byte) (key any, err error) {
	return x509.ParsePKCS8PrivateKey(der)
}
func (*Impl) ParsePKIXPublicKey(derBytes []byte) (pub any, err error) {
	return x509.ParsePKIXPublicKey(derBytes)
}
func (*Impl) SystemCertPool() (*x509.CertPool, error) {
	return x509.SystemCertPool()
}
