// Code generated by a tool. DO NOT EDIT.

// Package expvar provides a mockable wrapper for expvar.
package expvar

import (
	expvar "expvar"
	http "net/http"
)

var _ Interface = &Impl{}
var _ = expvar.Do

type Interface interface {
	Do(f func(expvar.KeyValue))
	Get(name string) expvar.Var
	Handler() http.Handler
	NewFloat(name string) *expvar.Float
	NewInt(name string) *expvar.Int
	NewMap(name string) *expvar.Map
	NewString(name string) *expvar.String
	Publish(name string, v expvar.Var)
}

type Impl struct{}

func (*Impl) Do(f func(expvar.KeyValue)) {
	expvar.Do(f)
}
func (*Impl) Get(name string) expvar.Var {
	return expvar.Get(name)
}
func (*Impl) Handler() http.Handler {
	return expvar.Handler()
}
func (*Impl) NewFloat(name string) *expvar.Float {
	return expvar.NewFloat(name)
}
func (*Impl) NewInt(name string) *expvar.Int {
	return expvar.NewInt(name)
}
func (*Impl) NewMap(name string) *expvar.Map {
	return expvar.NewMap(name)
}
func (*Impl) NewString(name string) *expvar.String {
	return expvar.NewString(name)
}
func (*Impl) Publish(name string, v expvar.Var) {
	expvar.Publish(name, v)
}
