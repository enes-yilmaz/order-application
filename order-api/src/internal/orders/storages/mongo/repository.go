package mongo

import (
	"OrderAPI/src/internal/orders/handlers/types"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
