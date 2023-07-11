package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/campbel/terpcatalog/admin"
	"github.com/campbel/terpcatalog/catalog"
	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
)

func main() {
	ctx := context.Background()

	catalogServer := catalog.NewServer(ctx, config.CatalogPort())
	go startServer("catalog", catalogServer)

	adminServer := admin.NewServer(ctx, config.AdminPort())
	go startServer("admin", adminServer)

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
