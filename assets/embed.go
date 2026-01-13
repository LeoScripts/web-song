package assets

import (
	"embed"
	"mime"
	"net/http"
)

//go:embed style.css
var assets embed.FS

var FS = http.FS(assets)

func init() {
	_ = mime.AddExtensionType(".css", "text/css")
}
