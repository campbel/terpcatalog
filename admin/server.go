package admin

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/campbel/terpcatalog/util/log"
)

var (
	//go:embed app/dist
	dist embed.FS
)

func NewServer(port string) *http.Server {
	fsys, err := fs.Sub(dist, "app/dist")
	if err != nil {
		log.FatalError("error during fs.Sub", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(fsys)))

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
