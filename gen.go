/*
	Package facades provides mockable interfaces for nearly every package in the standard library.

	Each package contains an Interface and a concrete Impl for its corresponding std package.
*/
package facades

//go:generate go run -modfile go-generate.mod jonwillia.ms/facade/cmd/facade -world -out .
