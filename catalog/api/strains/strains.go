package strains

import (
	"encoding/json"
	"net/http"

	"github.com/campbel/terpcatalog/shared/db/strains"
)

type Handler struct {
	strainsStore strains.Store
}

func NewHandler(strainsStore strains.Store) *Handler {
	return &Handler{
		strainsStore,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	strains, err := h.strainsStore.GetStrains(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(strains)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
