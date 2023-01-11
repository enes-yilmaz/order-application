package types

type (
	CreateOrderRequest struct {
		CustomerId string  `json:"customerId" bson:"customerId"`
		Quantity   int     `json:"quantity"`
		Price      float64 `json:"price"`
		Status     string  `json:"status"`
		Address    Address `json:"address"`
		Product    Product `json:"product"`
	}

	UpdateOrderRequest struct {
		CustomerId string  `json:"customerId" bson:"customerId"`
		Quantity   int     `json:"quantity"`
		Price      float64 `json:"price"`
		Status     string  `json:"status"`
		Address    Address `json:"address"`
		Product    Product `json:"product"`
	}
)

type AcceptOrderStatusType struct {
	Sku      string `json:"sku" validate:"required"`
	Source   string `json:"source" validate:"required"`
	OpenSale *bool  `json:"openSale"`
}

func (a AcceptOrderStatusType) AcceptableStatus() map[string]bool {
	return map[string]bool{
		"Prepared":  true,
		"Delivered": true,
		"OnRoad":    true,
		"Done":      true,
	}
}
