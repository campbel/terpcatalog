package admin

import (
	"context"
	"embed"
	"io/fs"
	"net/http"

	"github.com/campbel/terpcatalog/admin/api"
	"github.com/campbel/terpcatalog/admin/db/strains"
	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//go:embed app/dist/index.html
	index []byte

	//go:embed app/dist/favicon.ico
	favicon []byte

	//go:embed app/dist/assets
	assets embed.FS
)

func NewServer(ctx context.Context, port string) *http.Server {

	// Admin App Assets
	assetFS, err := fs.Sub(assets, "app/dist")
	if err != nil {
		log.FatalError("error during fs.Sub", err)
	}

	// Mongo Store
	clientOptions := options.Client().ApplyURI(config.MongoConnectionString())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	strainsCollection := client.Database("terpcatalog").Collection("strains")
	strainsStore := strains.NewStore(strainsCollection)

	mux := http.NewServeMux()
	mux.Handle("/", fileHandler("text/html", index))
	mux.Handle("/favicon.ico", fileHandler("image/x-icon", favicon))
	mux.Handle("/assets/", http.FileServer(http.FS(assetFS)))
	mux.Handle("/api/", api.NewHandler(strainsStore))

	return &http.Server{
		Addr:    ":" + port,
		Handler: logWrapper(mux),
	}
}

func fileHandler(contentType string, content []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write(content)
	})
}

func logWrapper(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", "method", r.Method, "path", r.URL.Path, "headers", r.Header)
		handler.ServeHTTP(w, r)
	})
}
