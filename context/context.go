// Code generated by a tool. DO NOT EDIT.

// Package context provides a mockable wrapper for context.
package context

import (
	context "context"
	time "time"
)

var _ Interface = &Impl{}
var _ = context.Background

type Interface interface {
	Background() context.Context
	TODO() context.Context
	WithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc)
	WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc)
	WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
	WithValue(parent context.Context, key any, val any) context.Context
}

type Impl struct{}

func (*Impl) Background() context.Context {
	return context.Background()
}
func (*Impl) TODO() context.Context {
	return context.TODO()
}
func (*Impl) WithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithCancel(parent)
}
func (*Impl) WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(parent, d)
}
func (*Impl) WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}
func (*Impl) WithValue(parent context.Context, key any, val any) context.Context {
	return context.WithValue(parent, key, val)
}
