package api

import (
	"net/http"

	sapi "github.com/campbel/terpcatalog/admin/api/strains"
	ss "github.com/campbel/terpcatalog/admin/db/strains"
)

func NewHandler(strainStore ss.Store) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/strains", sapi.NewHandler(strainStore))
	return mux
}
