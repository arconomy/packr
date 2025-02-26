package b

import "github.com/arconomy/packr"

func init() {
	packr.New("b-box", "../c")
	packr.New("cb-box", "../c")
}
