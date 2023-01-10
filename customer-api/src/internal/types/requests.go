package types

import (
	"time"
)

type (
	CustomerRequest struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Address   Address   `json:"address"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	GetAllCustomersRequest struct {
		Limit          int    `json:"limit"`
		Offset         int    `json:"offset"`
		OrderBy        string `json:"orderBy"`
		OrderDirection string `json:"orderDirection"`
		IsCount        bool   `json:"isCount"`
	}
	GetCustomerRequest struct {
		CustomerId string `json:"customerId"`
	}

	CreateCustomerRequest struct {
		Id      string  `json:"id" bson:"_id"`
		Name    string  `json:"name" validate:"required"`
		Email   string  `json:"email" validate:"required"`
		Address Address `json:"address"`
	}

	UpdateCustomerRequest struct {
		Id      string  `json:"id" bson:"_id" validate:"required"`
		Name    string  `json:"name"`
		Email   string  `json:"email"`
		Address Address `json:"address"`
	}
)
