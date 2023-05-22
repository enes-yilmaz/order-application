package mongo

import (
	"OrderAPI/src/internal/orders/handlers/types"
	"OrderAPI/src/pkg/errors"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{mc: mc}
}

func (r Repository) SaveOrder(order types.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := r.mc.InsertOne(ctx, order)
	return err
}

func (r Repository) UpdateOrder(order types.Order) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// filter for updated Order
	filter := bson.M{"_id": order.Id}

	update := bson.M{
		"$set": bson.M{
			"customerId": order.CustomerId,
			"quantity":   order.Quantity,
			"price":      order.Price,
			"status":     order.Status,
			"address":    order.Address,
			"product":    order.Product,
			"updatedAt":  order.UpdatedAt,
		},
	}

	// Update options
	opt := options.Update().SetUpsert(false)

	// Update document
	updatedResult, err := r.mc.UpdateOne(ctx, filter, update, opt)
	if err != nil {
		return 0, err
	}

	return int(updatedResult.ModifiedCount), nil
}

func (r Repository) DeleteOrderByOrderId(orderId string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{
		"_id": orderId,
	}
	res, err := r.mc.DeleteOne(ctx, filter)
	if err != nil {
		return 0, errors.DeleteOneFailed
	}

	return int(res.DeletedCount), nil
}

func (r Repository) ChangeOrderStatus(req types.ChangeOrderStatusRequest) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// filter for updated Order
	filter := bson.M{"_id": req.OrderId}

	update := bson.M{
		"$set": bson.M{
			"status":    req.Status,
			"updatedAt": time.Now(),
		},
	}
	opt := options.Update().SetUpsert(false)

	updatedResult, err := r.mc.UpdateOne(ctx, filter, update, opt)
	if err != nil {
		return 0, err
	}

	return int(updatedResult.ModifiedCount), nil
}

func (r Repository) GetAllOrders(req types.GetAllOrdersRequest) ([]types.Order, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	orders := make([]types.Order, 0)

	totalCount, err := r.mc.CountDocuments(ctx, bson.M{}, nil)
	if err != nil {
		return nil, 0, err
	}
	if req.IsCount {
		return nil, int(totalCount), nil
	}

	// Set findOptions
	findOptions := options.Find()
	if req.OrderBy == "" {
		req.OrderBy = "product.name"
	}
	if req.OrderDirection == "asc" || req.OrderDirection == "ASC" {
		findOptions.SetSort(bson.D{{req.OrderBy, 1}})
	} else {
		findOptions.SetSort(bson.D{{req.OrderBy, -1}})
	}
	findOptions.SetSkip(int64(req.Offset))
	findOptions.SetLimit(int64(req.Limit))

	cur, err := r.mc.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}

	if err := cur.All(ctx, &orders); err != nil {
		return nil, 0, err
	}
	return orders, int(totalCount), nil
}

func (r Repository) GetOrderByOrderId(req types.GetOrderByOrderIdRequest) (types.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var order types.Order
	filter := bson.M{
		"_id": req.OrderId,
	}

	err := r.mc.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return types.Order{}, errors.OrderNotFound
		}
		return types.Order{}, errors.FindFailed
	}

	return order, nil
}

func (r Repository) GetOrdersByCustomerId(req types.GetOrdersByCustomerIdRequest) ([]types.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	orders := make([]types.Order, 0)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"updatedAt", -1}})
	filter := bson.M{
		"customerId": req.CustomerId,
	}

	cur, err := r.mc.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}
