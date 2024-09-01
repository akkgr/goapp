package views

import (
	"embed"
	"io/fs"
)

//go:embed all:templates
var viewsFS embed.FS

func Templates() fs.FS {
	res, _ := fs.Sub(viewsFS, "templates")
	return res
}
