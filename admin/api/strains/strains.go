package strains

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/campbel/terpcatalog/util/log"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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

var strains []any

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(strains)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	var data any
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strains = append(strains, data)
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strains = append(strains[:id], strains[id+1:]...)
	w.WriteHeader(http.StatusNotFound)
}
