// Code generated by a tool. DO NOT EDIT.

// Package parse provides a mockable wrapper for text/template/parse.
package parse

import (
	parse "text/template/parse"
)

var _ Interface = &Impl{}
var _ = parse.IsEmptyTree

type Interface interface {
	IsEmptyTree(n parse.Node) bool
	New(name string, funcs ...map[string]any) *parse.Tree
	NewIdentifier(ident string) *parse.IdentifierNode
	Parse(name string, text string, leftDelim string, rightDelim string, funcs ...map[string]any) (map[string]*parse.Tree, error)
}

type Impl struct{}

func (*Impl) IsEmptyTree(n parse.Node) bool {
	return parse.IsEmptyTree(n)
}
func (*Impl) New(name string, funcs ...map[string]any) *parse.Tree {
	return parse.New(name, funcs...)
}
func (*Impl) NewIdentifier(ident string) *parse.IdentifierNode {
	return parse.NewIdentifier(ident)
}
func (*Impl) Parse(name string, text string, leftDelim string, rightDelim string, funcs ...map[string]any) (map[string]*parse.Tree, error) {
	return parse.Parse(name, text, leftDelim, rightDelim, funcs...)
}
