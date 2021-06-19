package frontend

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type fsFunc func(name string) (fs.File, error)

// Open implement fs.FS.
func (fs fsFunc) Open(name string) (fs.File, error) {
	return fs(name)
}

// assets filesystem,
//
// use: StaticFS(StaticPrefix, FileSystem(FS, StaticRoot))
// index.html page with string: IndexHtml
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

// Static returns a middleware handler that serves static files in the given directory.
func InitRouter(r *gin.Engine) *gin.Engine {
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
