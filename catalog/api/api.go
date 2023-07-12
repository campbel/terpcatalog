package api

import (
	"net/http"

	oapi "github.com/campbel/terpcatalog/catalog/api/orders"
	papi "github.com/campbel/terpcatalog/catalog/api/producers"
	sapi "github.com/campbel/terpcatalog/catalog/api/strains"
	os "github.com/campbel/terpcatalog/shared/db/orders"
	ps "github.com/campbel/terpcatalog/shared/db/producers"
	ss "github.com/campbel/terpcatalog/shared/db/strains"
)

func NewHandler(strainStore ss.Store, producerStore ps.Store, orderStore os.Store) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/strains", sapi.NewHandler(strainStore))
	mux.Handle("/api/producers", papi.NewHandler(producerStore))
	mux.Handle("/api/orders", oapi.NewHandler(orderStore))
	return mux
}
