package main

import (
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

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		done <- true
	}()

	go func() {
		log.Info("starting catalog server...", "port", config.Port())
		if err := http.ListenAndServe(":"+config.Port(), nil); err != nil {
			log.FatalError("error during http.ListenAndServe", err)
		}
		done <- true
	}()

	go func() {
		log.Info("starting admin server...", "port", config.AdminPort())
		adminServer := admin.NewServer(config.AdminPort())
		if err := adminServer.ListenAndServe(); err != nil {
			log.FatalError("error during adminServer.ListenAndServe", err)
		}
		done <- true
	}()

	<-done
	log.Info("shutting down...")
}
