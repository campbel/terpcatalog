package api

import (
	"net/http"

	papi "github.com/campbel/terpcatalog/admin/api/producers"
	sapi "github.com/campbel/terpcatalog/admin/api/strains"
	ps "github.com/campbel/terpcatalog/shared/db/producers"
	ss "github.com/campbel/terpcatalog/shared/db/strains"
)

func NewHandler(strainStore ss.Store, producerStore ps.Store) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/strains", sapi.NewHandler(strainStore))
	mux.Handle("/api/producers", papi.NewHandler(producerStore))
	return mux
}
