package packr

import "github.com/arconomy/packr/file"

func qfile(name string, body string) File {
	f, err := file.NewFile(name, []byte(body))
	if err != nil {
		panic(err)
	}
	return f
}
