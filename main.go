package main

import (
	"context"
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/campbel/terpcatalog/admin"
	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed public
	public embed.FS

	//go:embed data/catalog.yaml
	catalogYaml string
)

func main() {
	fsys, err := fs.Sub(public, "public")
	if err != nil {
		log.FatalError("error during fs.Sub", err)
	}

	var catalog any
	if err := yaml.Unmarshal([]byte(catalogYaml), &catalog); err != nil {
		log.FatalError("error during yaml.Unmarshal", err)
	}

	http.Handle("/api/catalog", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(catalog); err != nil {
			log.Error("error during json.Encode", err)
		}
	}))

	http.Handle("/", http.FileServer(http.FS(fsys)))

	catalogServer := &http.Server{
		Addr: ":" + config.Port(),
	}
	go startServer("catalog", catalogServer)

	adminServer := admin.NewServer(config.AdminPort())
	go startServer("admin", adminServer)

	ctx := context.Background()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	catalogServer.Shutdown(ctx)
	adminServer.Shutdown(ctx)
	log.Info("shutting down...")
}

func startServer(name string, server *http.Server) {
	log.Info("starting server...", "name", name, "addr", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("error during http.ListenAndServe", err)
	} else {
		log.Info("server stopped", "name", name)
	}
}
