package orders

import (
	"encoding/json"
	"net/http"

	"github.com/campbel/terpcatalog/shared/db/orders"
	"github.com/campbel/terpcatalog/shared/types"
)

type Handler struct {
	store orders.Store
}

func NewHandler(store orders.Store) *Handler {
	return &Handler{store}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	var data types.Order
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !isValidOrder(data) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order, err := h.store.CreateOrder(r.Context(), data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func isValidOrder(order types.Order) bool {
	if order.ID != "" {
		return false
	}
	if order.Information.Company == "" {
		return false
	}
	if order.Information.LicenseNumber == "" {
		return false
	}
	if order.Information.Email == "" {
		return false
	}
	if order.Information.Phone == "" {
		return false
	}
	if order.Information.Address.Street == "" {
		return false
	}
	if order.Information.Address.City == "" {
		return false
	}
	if order.Information.Address.State == "" {
		return false
	}
	if order.Information.Address.Postal == "" {
		return false
	}
	if len(order.Items) == 0 {
		return false
	}
	for _, item := range order.Items {
		if item.StrainID == "" {
			return false
		}
		if item.Quantity <= 0 {
			return false
		}
	}
	return true
}
