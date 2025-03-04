package store

const diskGlobalTmpl = `// +build !skippackr
// Code generated by github.com/arconomy/packr. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package {{.Package}}

import (
	"github.com/arconomy/packr"
	"github.com/arconomy/packr/file/resolver"
)

var _ = func() error {
	const gk = "{{.GK}}"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
	{{- range $k, $v := .GlobalFiles }}
		"{{$k}}": "{{$v}}",
	{{- end }}
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	{{- range $box := .Boxes}}
{{ printBox $box -}}
	{{ end }}
	return nil
}()
`

const diskImportTmpl = `// +build !skippackr
// Code generated by github.com/arconomy/packr. DO NOT EDIT.

// You can use the "packr clean" command to clean up this,
// and any other packr generated files.
package {{.Package}}

import _ "{{.Import}}"
`

const diskGlobalBoxTmpl = `
	func() {
		b := packr.New("{{.Box.Name}}", "{{.Box.Path}}")
		{{- range $file := .Files }}
		b.SetResolver("{{$file.Resolver}}", packr.Pointer{ForwardBox: gk, ForwardPath: "{{$file.ForwardPath}}"})
		{{- end }}
	}()`
