package dist

import (
	"embed"
	_ "embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

const (
	IndexHtml = "index.html"
	Prifix    = "/assets/"
)

//go:embed index.html assets
var Assets embed.FS

type fsFunc func(name string) (fs.File, error)

// Open implement fs.FS.
func (fs fsFunc) Open(name string) (fs.File, error) {
	return fs(name)
}

// AssetsHandler static assets handler.
//
// add to route:
// ServeMux.Handle("/prefix/", AssetsHandler("/prefix/", Assets, "./static"))
func AssetsHandler(prefix string, assets embed.FS, root string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetsPath := path.Join(root, name)
		file, err := assets.Open(assetsPath)
		if err != nil {
			return nil, err
		}
		return file, err
	})
	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}

// Static returns a middleware handler that serves static files in the given directory.
func GinHandlerFunc() gin.HandlerFunc {
	var mux = http.NewServeMux()
	mux.Handle(Prifix, AssetsHandler(Prifix, Assets, "assets"))
	return func(c *gin.Context) {
		mux.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func InitGinRouter(r *gin.Engine) *gin.Engine {
	t, err := template.ParseFS(Assets, IndexHtml)
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, IndexHtml, gin.H{"title": "Embed"})
	})
	r.Use(GinHandlerFunc())
	return r
}
