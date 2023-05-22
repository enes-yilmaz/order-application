package types

type (
	GetAllOrdersRequest struct {
		Limit          int    `json:"limit"`
		Offset         int    `json:"offset"`
		OrderBy        string `json:"orderBy"`
		OrderDirection string `json:"orderDirection"`
		IsCount        bool   `json:"isCount"`
	}

	GetOrderByOrderIdRequest struct {
		OrderId string `json:"orderId"`
	}

	GetOrdersByCustomerIdRequest struct {
		CustomerId string `json:"customerId"`
	}

	CreateOrderRequest struct {
		CustomerId string  `json:"customerId" bson:"customerId"`
		Quantity   int     `json:"quantity"`
		Price      float64 `json:"price"`
		Status     string  `json:"status"`
		Address    Address `json:"address"`
		Product    Product `json:"product"`
	}

	UpdateOrderRequest struct {
		// TODO: bson:"_id" olmali mi test et!!
		//OrderId    string  `json:"orderId" bson:"orderId"`
		OrderId    string  `json:"orderId" bson:"orderId"`
		CustomerId string  `json:"customerId" bson:"customerId"`
		Quantity   int     `json:"quantity"`
		Price      float64 `json:"price"`
		Status     string  `json:"status"`
		Address    Address `json:"address"`
		Product    Product `json:"product"`
	}

	ChangeOrderStatusRequest struct {
		OrderId string `json:"orderId" bson:"_id"`
		Status  string `json:"status"`
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
