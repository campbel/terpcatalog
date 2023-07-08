package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
)

//go:embed public
var public embed.FS

func main() {
	log.Info("starting server", "port", config.Port())
	fsys, err := fs.Sub(public, "public")
	if err != nil {
		log.FatalError("error during fs.Sub", err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))
	if err := http.ListenAndServe(":"+config.Port(), nil); err != nil {
		log.FatalError("error during http.ListenAndServe", err)
	}
}
