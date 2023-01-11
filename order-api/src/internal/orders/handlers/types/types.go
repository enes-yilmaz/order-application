package types

import "time"

type (
	Order struct {
		Id         string    `json:"id" bson:"_id"`
		CustomerId string    `json:"customerId" bson:"customerId"`
		Quantity   int       `json:"quantity"`
		Price      float64   `json:"price"`
		Status     string    `json:"status"`
		Address    Address   `json:"address"`
		Product    Product   `json:"product"`
		CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
	}

	Address struct {
		AddressLine string `json:"addressLine" bson:"addressLine"`
		City        string `json:"city"`
		Country     string `json:"country"`
		CityCode    int    `json:"cityCode" bson:"cityCode"`
	}

	Product struct {
		Id       string `json:"id" bson:"id"`
		ImageUrl string `json:"imageUrl" bson:"imageUrl"`
		Name     string `json:"name"`
	}
)
