package admin

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/campbel/terpcatalog/admin/api"
	"github.com/campbel/terpcatalog/util/log"
)

var (
	//go:embed app/dist
	dist embed.FS
)

func NewServer(port string) *http.Server {

	// Admin App
	fsys, err := fs.Sub(dist, "app/dist")
	if err != nil {
		log.FatalError("error during fs.Sub", err)
	}

	// Admin API
	handler := api.NewHandler()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(fsys)))
	mux.Handle("/api/", handler)

	return &http.Server{
		Addr:    ":" + port,
		Handler: logWrapper(mux),
	}
}

func logWrapper(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", "method", r.Method, "path", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
