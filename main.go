package main

import (
	"net/http"

	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
)

func main() {
	log.Info("starting server", "port", config.Port())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	if err := http.ListenAndServe(":"+config.Port(), nil); err != nil {
		log.FatalError("error during http.ListenAndServe", err)
	}
}
