// Code generated by a tool. DO NOT EDIT.

// Package constant provides a mockable wrapper for go/constant.
package constant

import (
	constant "go/constant"
	token "go/token"
)

var _ Interface = &Impl{}
var _ = constant.BinaryOp

type Interface interface {
	BinaryOp(x_ constant.Value, op token.Token, y_ constant.Value) constant.Value
	BitLen(x constant.Value) int
	BoolVal(x constant.Value) bool
	Bytes(x constant.Value) []byte
	Compare(x_ constant.Value, op token.Token, y_ constant.Value) bool
	Denom(x constant.Value) constant.Value
	Float32Val(x constant.Value) (float32, bool)
	Float64Val(x constant.Value) (float64, bool)
	Imag(x constant.Value) constant.Value
	Int64Val(x constant.Value) (int64, bool)
	Make(x any) constant.Value
	MakeBool(b bool) constant.Value
	MakeFloat64(x float64) constant.Value
	MakeFromBytes(bytes []byte) constant.Value
	MakeFromLiteral(lit string, tok token.Token, zero uint) constant.Value
	MakeImag(x constant.Value) constant.Value
	MakeInt64(x int64) constant.Value
	MakeString(s string) constant.Value
	MakeUint64(x uint64) constant.Value
	MakeUnknown() constant.Value
	Num(x constant.Value) constant.Value
	Real(x constant.Value) constant.Value
	Shift(x constant.Value, op token.Token, s uint) constant.Value
	Sign(x constant.Value) int
	StringVal(x constant.Value) string
	ToComplex(x constant.Value) constant.Value
	ToFloat(x constant.Value) constant.Value
	ToInt(x constant.Value) constant.Value
	Uint64Val(x constant.Value) (uint64, bool)
	UnaryOp(op token.Token, y constant.Value, prec uint) constant.Value
	Val(x constant.Value) any
}

type Impl struct{}

func (*Impl) BinaryOp(x_ constant.Value, op token.Token, y_ constant.Value) constant.Value {
	return constant.BinaryOp(x_, op, y_)
}
func (*Impl) BitLen(x constant.Value) int {
	return constant.BitLen(x)
}
func (*Impl) BoolVal(x constant.Value) bool {
	return constant.BoolVal(x)
}
func (*Impl) Bytes(x constant.Value) []byte {
	return constant.Bytes(x)
}
func (*Impl) Compare(x_ constant.Value, op token.Token, y_ constant.Value) bool {
	return constant.Compare(x_, op, y_)
}
func (*Impl) Denom(x constant.Value) constant.Value {
	return constant.Denom(x)
}
func (*Impl) Float32Val(x constant.Value) (float32, bool) {
	return constant.Float32Val(x)
}
func (*Impl) Float64Val(x constant.Value) (float64, bool) {
	return constant.Float64Val(x)
}
func (*Impl) Imag(x constant.Value) constant.Value {
	return constant.Imag(x)
}
func (*Impl) Int64Val(x constant.Value) (int64, bool) {
	return constant.Int64Val(x)
}
func (*Impl) Make(x any) constant.Value {
	return constant.Make(x)
}
func (*Impl) MakeBool(b bool) constant.Value {
	return constant.MakeBool(b)
}
func (*Impl) MakeFloat64(x float64) constant.Value {
	return constant.MakeFloat64(x)
}
func (*Impl) MakeFromBytes(bytes []byte) constant.Value {
	return constant.MakeFromBytes(bytes)
}
func (*Impl) MakeFromLiteral(lit string, tok token.Token, zero uint) constant.Value {
	return constant.MakeFromLiteral(lit, tok, zero)
}
func (*Impl) MakeImag(x constant.Value) constant.Value {
	return constant.MakeImag(x)
}
func (*Impl) MakeInt64(x int64) constant.Value {
	return constant.MakeInt64(x)
}
func (*Impl) MakeString(s string) constant.Value {
	return constant.MakeString(s)
}
func (*Impl) MakeUint64(x uint64) constant.Value {
	return constant.MakeUint64(x)
}
func (*Impl) MakeUnknown() constant.Value {
	return constant.MakeUnknown()
}
func (*Impl) Num(x constant.Value) constant.Value {
	return constant.Num(x)
}
func (*Impl) Real(x constant.Value) constant.Value {
	return constant.Real(x)
}
func (*Impl) Shift(x constant.Value, op token.Token, s uint) constant.Value {
	return constant.Shift(x, op, s)
}
func (*Impl) Sign(x constant.Value) int {
	return constant.Sign(x)
}
func (*Impl) StringVal(x constant.Value) string {
	return constant.StringVal(x)
}
func (*Impl) ToComplex(x constant.Value) constant.Value {
	return constant.ToComplex(x)
}
func (*Impl) ToFloat(x constant.Value) constant.Value {
	return constant.ToFloat(x)
}
func (*Impl) ToInt(x constant.Value) constant.Value {
	return constant.ToInt(x)
}
func (*Impl) Uint64Val(x constant.Value) (uint64, bool) {
	return constant.Uint64Val(x)
}
func (*Impl) UnaryOp(op token.Token, y constant.Value, prec uint) constant.Value {
	return constant.UnaryOp(op, y, prec)
}
func (*Impl) Val(x constant.Value) any {
	return constant.Val(x)
}