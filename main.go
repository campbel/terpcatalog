package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
)

//go:embed public
var public embed.FS

func main() {
	log.Info("starting server...", "port", config.Port())
	fsys, err := fs.Sub(public, "public")
	if err != nil {
		log.FatalError("error during fs.Sub", err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		done <- true
	}()

	go func() {
		if err := http.ListenAndServe(":"+config.Port(), nil); err != nil {
			log.FatalError("error during http.ListenAndServe", err)
		}
		done <- true
	}()

	<-done
	log.Info("shutting down...")
}
