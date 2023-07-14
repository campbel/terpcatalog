package strains

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/campbel/terpcatalog/shared/db/strains"
	"github.com/campbel/terpcatalog/shared/types"
	"github.com/campbel/terpcatalog/util/log"
)

type Handler struct {
	store strains.Store
}

func NewHandler(store strains.Store) *Handler {
	return &Handler{store}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		log.Debug("error decoding strain data", "error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := isValid(data); err != nil {
		log.Debug("error in strain data", "error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	strain, err := h.store.CreateStrain(r.Context(), data)
	if err != nil {
		log.Debug("error creating strain in store", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(strain)
	if err != nil {
		log.Debug("error encoding strain data", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func isValid(data types.Strain) error {
	if data.ProducerID == "" {
		return errors.New("producer id is required")
	}
	if data.Name == "" {
		return errors.New("name is required")
	}
	if data.Category == "" {
		return errors.New("category is required")
	}
	if data.Genetics == "" {
		return errors.New("genetics is required")
	}
	if data.THC < 0 || data.THC > 100 {
		return errors.New("thc must be between 0 and 100")
	}
	if data.CBD < 0 || data.CBD > 100 {
		return errors.New("cbd must be between 0 and 100")
	}
	if data.Terpenes < 0 || data.Terpenes > 100 {
		return errors.New("terpenes must be between 0 and 100")
	}
	if data.TotalCannabinoids < 0 || data.TotalCannabinoids > 100 {
		return errors.New("total cannabinoids must be between 0 and 100")
	}
	for _, terp := range data.TerpeneList {
		if terp == "" {
			return errors.New("terpene list cannot contain empty strings")
		}
	}
	if data.Price < 0 {
		return errors.New("price cannot be negative")
	}
	if data.HarvestDate == "" {
		return errors.New("harvest date is required")
	}
	for _, img := range data.Images {
		if img == "" {
			return errors.New("image list cannot contain empty strings")
		}
	}
	return nil
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
