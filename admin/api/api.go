package api

import (
	"net/http"

	"github.com/campbel/terpcatalog/admin/api/strains"
)

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/strains", strains.NewHandler())
	return mux
}
