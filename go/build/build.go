// Code generated by a tool. DO NOT EDIT.

// Package build provides a mockable wrapper for go/build.
package build

import (
	build "go/build"
)

var _ Interface = &Impl{}
var _ = build.ArchChar

type Interface interface {
	ArchChar(goarch string) (string, error)
	Import(path string, srcDir string, mode build.ImportMode) (*build.Package, error)
	ImportDir(dir string, mode build.ImportMode) (*build.Package, error)
	IsLocalImport(path string) bool
}

type Impl struct{}

func (*Impl) ArchChar(goarch string) (string, error) {
	return build.ArchChar(goarch)
}
func (*Impl) Import(path string, srcDir string, mode build.ImportMode) (*build.Package, error) {
	return build.Import(path, srcDir, mode)
}
func (*Impl) ImportDir(dir string, mode build.ImportMode) (*build.Package, error) {
	return build.ImportDir(dir, mode)
}
func (*Impl) IsLocalImport(path string) bool {
	return build.IsLocalImport(path)
}