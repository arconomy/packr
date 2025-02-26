package store

import (
	"github.com/arconomy/packr/jam/parser"
)

type Store interface {
	FileNames(*parser.Box) ([]string, error)
	Files(*parser.Box) ([]*parser.File, error)
	Pack(*parser.Box) error
	Clean(*parser.Box) error
}
