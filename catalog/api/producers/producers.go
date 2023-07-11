package producers

import (
	"encoding/json"
	"net/http"

	"github.com/campbel/terpcatalog/shared/db/producers"
	"github.com/campbel/terpcatalog/util/log"
)

type Handler struct {
	store producers.Store
}

func NewHandler(store producers.Store) *Handler {
	return &Handler{store}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("strain api", "method", r.Method, "path", r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	producers, err := h.store.GetProducers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(producers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
