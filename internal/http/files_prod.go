//go:build !dev

package http

import (
	"embed"
)

//go:embed web
var fsys embed.FS
