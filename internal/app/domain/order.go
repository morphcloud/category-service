package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

//easyjson:json
type OrderList []Order

type Order struct {
	ID           string              `json:"id"           bson:"_id"`
	RestaurantID int64               `json:"restaurantId" bson:"restaurantId"`
	ConsumerID   int64               `json:"consumerId"   bson:"consumerId"`
	DeliveryID   int64               `json:"deliveryId"   bson:"deliveryId"`
	PaymentID    int64               `json:"paymentId"    bson:"paymentId"`
	StatusID     int32               `json:"statusId"     bson:"statusId"`
	CreatedAt    primitive.Timestamp `json:"createdAt"    bson:"createdAt"`
	UpdatedAt    primitive.Timestamp `json:"updatedAt"    bson:"updatedAt"`
}

type LineItem struct {
	ID         string  `json:"id"         bson:"_id"`
	MenuItemID int64   `json:"menuItemId" bson:"menuItemId"`
	OrderID    int64   `json:"orderId"    bson:"orderId"`
	Name       string  `json:"name"       bson:"name"`
	Price      float32 `json:"price"      bson:"price"`
	Quantity   int32   `json:"quantity"   bson:"quantity"`
}
