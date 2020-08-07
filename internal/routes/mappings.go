package routes

import (
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/morphcloud/order-service/internal/app/handlers/http/rest/v1"
	"github.com/morphcloud/order-service/internal/app/repositories"
	"github.com/morphcloud/order-service/internal/app/services"
	"github.com/morphcloud/order-service/internal/diagnostics"
)

func MapURLPathsToHandlers(r *mux.Router, c *mongo.Client, l *log.Logger) {
	r.HandleFunc("/healthz", diagnostics.LivenessHandler(l))
	r.HandleFunc("/readyz", diagnostics.ReadinessHandler(l))

	orderRepository := repositories.NewMongoOrderRepository(c)
	orderService := services.NewOrderService(orderRepository)
	orderHandler := v1.NewOrderHandler(orderService)

	//r.HandleFunc("/v1/orders", orderHandler.FindAll).Methods("GET")
	//r.Methods("GET").Path("/v1/orders").HandlerFunc(orderHandler.FindAll)
	v1Orders := r.PathPrefix("/v1/orders").Subrouter()
	v1Orders.Methods("GET").Path("").HandlerFunc(orderHandler.FindAll)
	v1Orders.Methods("GET").Path("/{order_id}").HandlerFunc(orderHandler.FindOne)
}
