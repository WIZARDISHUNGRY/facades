// Code generated by a tool. DO NOT EDIT.

// Package httptrace provides a mockable wrapper for net/http/httptrace.
package httptrace

import (
	context "context"
	httptrace "net/http/httptrace"
)

var _ Interface = &Impl{}
var _ = httptrace.ContextClientTrace

type Interface interface {
	ContextClientTrace(ctx context.Context) *httptrace.ClientTrace
	WithClientTrace(ctx context.Context, trace *httptrace.ClientTrace) context.Context
}

type Impl struct{}

func (*Impl) ContextClientTrace(ctx context.Context) *httptrace.ClientTrace {
	return httptrace.ContextClientTrace(ctx)
}
func (*Impl) WithClientTrace(ctx context.Context, trace *httptrace.ClientTrace) context.Context {
	return httptrace.WithClientTrace(ctx, trace)
}