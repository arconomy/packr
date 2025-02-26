package resolver

import "github.com/arconomy/packr/file"

type Packable interface {
	Pack(name string, f file.File) error
}
