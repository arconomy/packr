package import_pkg

import (
	"github.com/arconomy/packr"
)

var BoxTestNew = packr.New("pkg_test", "./pkg_test")
var BoxTestNewBox = packr.NewBox("./pkg_test")
