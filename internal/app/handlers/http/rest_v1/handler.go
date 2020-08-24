package rest_v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"github.com/morphcloud/order-service/internal/app/services"
	"github.com/morphcloud/order-service/pkg/http_response"
)

type OrderHandler interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindOne(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) OrderHandler {
	return &orderHandler{
		service,
	}
}

func (h *orderHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := http_response.MultipleJSONResponse{Data: orders}
	if _, _, err = easyjson.MarshalToHTTPResponseWriter(resp, w); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *orderHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["order_id"]
	order, err := h.service.FindOne(r.Context(), orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if _, _, err = easyjson.MarshalToHTTPResponseWriter(order, w); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *orderHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func (h *orderHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func (h *orderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
