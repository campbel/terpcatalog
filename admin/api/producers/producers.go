package producers

import (
	"encoding/json"
	"net/http"

	"github.com/campbel/terpcatalog/shared/db/producers"
	"github.com/campbel/terpcatalog/shared/types"
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
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodDelete:
		h.handleDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		strain, err := h.store.GetProducer(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode([]types.Producer{strain})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
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

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	var data types.Producer
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strain, err := h.store.CreateProducer(r.Context(), data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(strain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.store.DeleteProducer(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
