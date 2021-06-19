package frontend

import (
	"embed"
	_ "embed"
)

const (
	RootDir      = "dist"
	StaticPrefix = "/assets/"
	StaticRoot   = RootDir + StaticPrefix
	IndexHtml    = "index.html"
	Patterns     = RootDir + "/*.html"
)

//go:embed dist
var FS embed.FS
