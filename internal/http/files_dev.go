//go:build dev

package http

import (
	"io/fs"
	"os"
)

var fsys fs.FS = os.DirFS("./")
