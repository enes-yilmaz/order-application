package mongo

import (
	types2 "CustomerAPI/src/internal/storages/mongo/types"
	"CustomerAPI/src/internal/types"
	"CustomerAPI/src/pkg/errors"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{mc: mc}
}

func (r Repository) GetAllCustomers(req types.GetAllCustomersRequest) (types.Customers, error) {

	findCustomersFilter := types2.FindCustomersFilter{
		Limit:   int64(req.Limit),
		Offset:  int64(req.Offset),
		IsCount: req.IsCount,
	}
	if req.OrderBy == "" {
		req.OrderBy = "name"
	}
	findCustomersFilter.OrderBy = req.OrderBy

	if req.OrderDirection == "asc" || req.OrderDirection == "ASC" {
		findCustomersFilter.OrderDirection = 1
	} else {
		findCustomersFilter.OrderDirection = -1
	}

	customers, err := r.findCustomersByAggregation(findCustomersFilter.BuildCustomersQueryPipeline())
	if err != nil {
		return types.Customers{}, errors.FindFailed
	}

	return customers, nil

}

func (r Repository) GetCustomer(req types.GetCustomerRequest) (types.Customer, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var customer types.Customer
	filter := bson.M{
		"_id": req.CustomerId,
	}

	err := r.mc.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return types.Customer{}, errors.CustomerNotFound
		}
		return types.Customer{}, errors.FindFailed
	}

	return customer, nil
}

func (r Repository) IsCustomerExist(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var customer types.Customer
	filter := bson.M{
		"email": email,
	}
	err := r.mc.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		return false
	}

	return true
}

func (r Repository) IsValidCustomer(req types.GetCustomerRequest) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var customer types.Customer
	filter := bson.M{
		"_id": req.CustomerId,
	}

	err := r.mc.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.FindFailed
	}

	return true, nil
}

func (r Repository) SaveCustomer(customer types.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := r.mc.InsertOne(ctx, customer)
	return err
}

func (r Repository) DeleteCustomerById(customerId string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{
		"_id": customerId,
	}
	res, err := r.mc.DeleteOne(ctx, filter)
	if err != nil {
		return 0, errors.DeleteOneFailed
	}

	return res.DeletedCount, nil
}

func (r Repository) UpdateOneByCustomerId2(filter, updateModel bson.M) (types.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var customer types.Customer

	res, err := r.mc.UpdateOne(ctx, filter, updateModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return types.Customer{}, errors.CustomerNotFound
		}
		return types.Customer{}, errors.UpdateOneFailed
	}

	if res.MatchedCount == 0 {
		return types.Customer{}, errors.CustomerNotFound
	}
	_ = r.mc.FindOne(ctx, filter).Decode(&customer)

	return customer, nil
}

func (r Repository) UpdateOneByCustomerId(filter, updateModel bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := r.mc.UpdateOne(ctx, filter, updateModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.CustomerNotFound
		}
		return errors.UpdateOneFailed
	}

	if res.MatchedCount == 0 {
		return errors.CustomerNotFound
	}

	return nil
}

func (r Repository) findCustomersByAggregation(pipeline []bson.M) (types.Customers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	customers := make([]types.Customers, 0)
	cur, err := r.mc.Aggregate(ctx, pipeline)

	if err != nil {
		return customers[0], errors.FindFailed
	}

	if err := cur.All(ctx, &customers); err != nil {
		return customers[0], errors.MongoCursorFailed
	}

	return customers[0], nil
}
