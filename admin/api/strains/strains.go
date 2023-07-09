package strains

import (
	"encoding/json"
	"net/http"

	"github.com/campbel/terpcatalog/admin/db/strains"
	"github.com/campbel/terpcatalog/admin/types"
	"github.com/campbel/terpcatalog/util/log"
)

type Handler struct {
	store strains.Store
}

func NewHandler(store strains.Store) *Handler {
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
		strain, err := h.store.GetStrain(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode([]types.Strain{strain})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	strains, err := h.store.GetStrains(r.Context())
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

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	var data types.Strain
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strain, err := h.store.CreateStrain(r.Context(), data)
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
	err := h.store.DeleteStrain(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
