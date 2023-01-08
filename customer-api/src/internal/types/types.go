package types

import "time"

type Customer struct {
	Id        string    `json:"id" bson:"_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Address struct {
	AddressLine string `json:"addressLine" bson:"addressLine"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CityCode    int    `json:"cityCode" bson:"cityCode"`
}
