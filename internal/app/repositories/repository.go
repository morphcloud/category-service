package repositories

import (
	"context"

	"github.com/morphcloud/order-service/internal/app/domain"
)

// OrderRepository
type OrderRepository interface {
	FindAll(ctx context.Context) (domain.OrderList, error)
	FindOne(ctx context.Context, orderID string) (domain.Order, error)
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, order domain.Order) (domain.Order, error)
	Delete(ctx context.Context, orderID string) (bool, error)
}
