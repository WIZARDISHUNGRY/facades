// Code generated by a tool. DO NOT EDIT.

// Package importer provides a mockable wrapper for go/importer.
package importer

import (
	importer "go/importer"
	token "go/token"
	types "go/types"
)

var _ Interface = &Impl{}
var _ = importer.Default

type Interface interface {
	Default() types.Importer
	For(compiler string, lookup importer.Lookup) types.Importer
	ForCompiler(fset *token.FileSet, compiler string, lookup importer.Lookup) types.Importer
}

type Impl struct{}

func (*Impl) Default() types.Importer {
	return importer.Default()
}
func (*Impl) For(compiler string, lookup importer.Lookup) types.Importer {
	return importer.For(compiler, lookup)
}
func (*Impl) ForCompiler(fset *token.FileSet, compiler string, lookup importer.Lookup) types.Importer {
	return importer.ForCompiler(fset, compiler, lookup)
}
