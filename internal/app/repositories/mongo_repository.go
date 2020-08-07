package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/morphcloud/order-service/internal/app/domain"
)

type mongoOrderRepository struct {
	client *mongo.Client
}

// NewMongoOrderRepository
func NewMongoOrderRepository(client *mongo.Client) OrderRepository {
	return &mongoOrderRepository{client}
}

// FindAll
func (r *mongoOrderRepository) FindAll(ctx context.Context) (domain.OrderList, error) {
	cursor, err := r.client.Database("orderDB").Collection("orders").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders domain.OrderList

	for cursor.Next(ctx) {
		var order domain.Order
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// FindOne
func (r *mongoOrderRepository) FindOne(ctx context.Context, orderID string) (domain.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return domain.Order{}, err
	}

	res := r.client.
		Database("orderDB").
		Collection("orders").
		FindOne(ctx, bson.D{{"_id", objectID}})

	var order domain.Order
	if err := res.Decode(&order); err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

// Create TODO
func (r *mongoOrderRepository) Create(ctx context.Context, order domain.Order) (domain.Order, error) {
	return domain.Order{}, nil
}

// Update TODO
func (r *mongoOrderRepository) Update(ctx context.Context, order domain.Order) (domain.Order, error) {
	return domain.Order{}, nil
}

// Delete TODO
func (r *mongoOrderRepository) Delete(ctx context.Context, orderID string) (bool, error) {
	return false, nil
}
