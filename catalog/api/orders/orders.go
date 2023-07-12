package orders

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/campbel/terpcatalog/shared/db/orders"
	"github.com/campbel/terpcatalog/shared/email"
	"github.com/campbel/terpcatalog/shared/types"
	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
)

type Handler struct {
	store  orders.Store
	sender email.Sender
}

func NewHandler(store orders.Store, sender email.Sender) *Handler {
	return &Handler{store, sender}
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

	// Don't send e-mails in development
	if !config.IsDevelopment() {
		content, err := generateConfirmationTemplate(order)
		if err != nil {
			log.Error("error getting order email content", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = h.sender.SendEmail([]types.EmailAddress{{
			Name:  order.Information.Company,
			Email: order.Information.Email,
		}},
			"TerpScout Order Confirmation", content)
		if err != nil {
			log.Error("error sending email content", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content, err = generateOrderTemplate(order)
		if err != nil {
			log.Error("error getting order email content", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = h.sender.SendEmail(config.EmailOrderList(), "TerpScout New Order", content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
		if item.Strain.ID == "" {
			return false
		}
		if item.Quantity <= 0 {
			return false
		}
	}
	return true
}

func renderTemplate(t *template.Template, data any) (string, error) {
	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
