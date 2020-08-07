package services

import (
	"context"

	"github.com/morphcloud/order-service/internal/app/domain"
	"github.com/morphcloud/order-service/internal/app/repositories"
)

// OrderService
type OrderService interface {
	FindAll(ctx context.Context) (domain.OrderList, error)
	FindOne(ctx context.Context, orderID string) (domain.Order, error)
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, order domain.Order) (domain.Order, error)
	Delete(ctx context.Context, orderID string) (bool, error)
}

type orderService struct {
	repository repositories.OrderRepository
}

// NewOrderService
func NewOrderService(repository repositories.OrderRepository) OrderService {
	return &orderService{
		repository,
	}
}

func (s *orderService) FindAll(ctx context.Context) (domain.OrderList, error) {
	return s.repository.FindAll(ctx)
}

func (s *orderService) FindOne(ctx context.Context, orderID string) (domain.Order, error) {
	return s.repository.FindOne(ctx, orderID)
}

func (s *orderService) Create(ctx context.Context, order domain.Order) (domain.Order, error) {
	return s.repository.Create(ctx, order)
}

func (s *orderService) Update(ctx context.Context, order domain.Order) (domain.Order, error) {
	return s.repository.Update(ctx, order)
}

func (s *orderService) Delete(ctx context.Context, orderID string) (bool, error) {
	return s.repository.Delete(ctx, orderID)
}
