package admin

import (
	"context"
	"embed"
	"io/fs"
	"net/http"

	"github.com/campbel/terpcatalog/util/log"
)

var (
	//go:embed app/dist
	dist embed.FS
)

type Server struct {
	server http.Server
}

func NewServer(port string) *Server {
	fsys, err := fs.Sub(dist, "app/dist")
	if err != nil {
		log.FatalError("error during fs.Sub", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(fsys)))

	return &Server{
		server: http.Server{
			Addr:    ":" + port,
			Handler: mux,
		},
	}
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
