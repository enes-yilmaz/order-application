package clients

import (
	cfg "OrderAPI/src/config"
	"OrderAPI/src/internal/orders/clients/types"
	"OrderAPI/src/internal/orders/errors"
	rest "OrderAPI/src/pkg/clients"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type Customer struct {
	client *rest.BaseClient
}

func NewCustomerClient() *Customer {
	return &Customer{
		client: rest.NewBaseClient(cfg.GetConfigs().CustomerClient),
	}
}

func (c Customer) ValidateCustomer(customerId string) (bool, error) {
	var response types.ValidCustomerResponse

	res, err := c.client.GET("/validate/" + customerId)
	if err != nil {
		return response.IsValidCustomer, errors.ValidateCustomerError
	}

	defer fasthttp.ReleaseResponse(res)

	if err = json.Unmarshal(res.Body(), &response); err != nil {
		return response.IsValidCustomer, errors.FailedToBindError
	}

	return response.IsValidCustomer, nil
}
