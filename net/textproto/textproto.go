// Code generated by a tool. DO NOT EDIT.

// Package textproto provides a mockable wrapper for net/textproto.
package textproto

import (
	bufio "bufio"
	io "io"
	textproto "net/textproto"
)

var _ Interface = &Impl{}
var _ = textproto.CanonicalMIMEHeaderKey

type Interface interface {
	CanonicalMIMEHeaderKey(s string) string
	Dial(network string, addr string) (*textproto.Conn, error)
	NewConn(conn io.ReadWriteCloser) *textproto.Conn
	NewReader(r *bufio.Reader) *textproto.Reader
	NewWriter(w *bufio.Writer) *textproto.Writer
	TrimBytes(b []byte) []byte
	TrimString(s string) string
}

type Impl struct{}

func (*Impl) CanonicalMIMEHeaderKey(s string) string {
	return textproto.CanonicalMIMEHeaderKey(s)
}
func (*Impl) Dial(network string, addr string) (*textproto.Conn, error) {
	return textproto.Dial(network, addr)
}
func (*Impl) NewConn(conn io.ReadWriteCloser) *textproto.Conn {
	return textproto.NewConn(conn)
}
func (*Impl) NewReader(r *bufio.Reader) *textproto.Reader {
	return textproto.NewReader(r)
}
func (*Impl) NewWriter(w *bufio.Writer) *textproto.Writer {
	return textproto.NewWriter(w)
}
func (*Impl) TrimBytes(b []byte) []byte {
	return textproto.TrimBytes(b)
}
func (*Impl) TrimString(s string) string {
	return textproto.TrimString(s)
}