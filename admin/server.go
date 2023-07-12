package admin

import (
	"context"
	"embed"
	"io/fs"
	"net/http"

	"github.com/campbel/terpcatalog/admin/api"
	"github.com/campbel/terpcatalog/shared/db/producers"
	"github.com/campbel/terpcatalog/shared/db/strains"
	inthttp "github.com/campbel/terpcatalog/shared/http"
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

	mux := http.NewServeMux()
	mux.Handle("/", inthttp.FileHandler("text/html", index))
	mux.Handle("/favicon.ico", inthttp.FileHandler("image/x-icon", favicon))
	mux.Handle("/assets/", http.FileServer(http.FS(assetFS)))
	mux.Handle("/api/", api.NewHandler(
		strains.NewStore(client.Database("terpcatalog").Collection("strains")),
		producers.NewStore(client.Database("terpcatalog").Collection("producers")),
	))

	return &http.Server{
		Addr:    ":" + port,
		Handler: inthttp.LogWrapper(mux),
	}
}
