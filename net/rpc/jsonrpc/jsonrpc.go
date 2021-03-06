// Code generated by a tool. DO NOT EDIT.

// Package jsonrpc provides a mockable wrapper for net/rpc/jsonrpc.
package jsonrpc

import (
	io "io"
	rpc "net/rpc"
	jsonrpc "net/rpc/jsonrpc"
)

var _ Interface = &Impl{}
var _ = jsonrpc.Dial

type Interface interface {
	Dial(network string, address string) (*rpc.Client, error)
	NewClient(conn io.ReadWriteCloser) *rpc.Client
	NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec
	NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec
	ServeConn(conn io.ReadWriteCloser)
}

type Impl struct{}

func (*Impl) Dial(network string, address string) (*rpc.Client, error) {
	return jsonrpc.Dial(network, address)
}
func (*Impl) NewClient(conn io.ReadWriteCloser) *rpc.Client {
	return jsonrpc.NewClient(conn)
}
func (*Impl) NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec {
	return jsonrpc.NewClientCodec(conn)
}
func (*Impl) NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return jsonrpc.NewServerCodec(conn)
}
func (*Impl) ServeConn(conn io.ReadWriteCloser) {
	jsonrpc.ServeConn(conn)
}
