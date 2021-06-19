package frontend

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func init() {
}

type fsFunc func(name string) (fs.File, error)

// Open implement fs.FS.
func (fs fsFunc) Open(name string) (fs.File, error) {
	return fs(name)
}

func FileSystem(assets embed.FS, root string) http.FileSystem {
	return http.FS(fsFunc(func(name string) (fs.File, error) {
		file, err := assets.Open(path.Join(root, name))
		if err != nil {
			return nil, err
		}
		file.Close()
		return file, err
	}))
}

// Handler static assets handler.
func Handler() http.Handler {
	return http.FileServer(FileSystem(FS, StaticRoot))
}

func Static(prefix string) {
	http.StripPrefix(prefix, Handler())
}

// Static returns a middleware handler that serves static files in the given directory.
func InitGinRouter(r *gin.Engine) *gin.Engine {
	// assets file
	r.StaticFS(StaticPrefix, FileSystem(FS, StaticRoot))

	// fe, err := fs.Sub(FS, "dist")
	// if err != nil {
	// 	panic(err)
	// }
	// tmpl := template.Must(template.New("").ParseFS(fe, "*.html"))

	tmpl := template.Must(template.New("").ParseFS(FS, Patterns))
	r.SetHTMLTemplate(tmpl)
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, IndexHtml, gin.H{"title": "Index"})
	})
	return r
}
