package resolver

import "github.com/arconomy/packr/file"

func qfile(name string, body string) file.File {
	f, err := file.NewFile(name, []byte(body))
	if err != nil {
		panic(err)
	}
	return f
}
