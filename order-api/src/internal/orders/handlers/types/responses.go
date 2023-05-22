package types

type (
	TotalCount struct {
		Count int `bson:"count" json:"count"`
	}
	Orders struct {
		Items      []Order    `bson:"items" json:"items"`
		TotalCount TotalCount `bson:"totalCount" json:"totalCount"`
	}
	GetOrdersResponse struct {
		Items      []Order `json:"items"`
		TotalCount int     `json:"totalCount"`
		Limit      int     `json:"limit"`
		Offset     int     `json:"offset"`
	}
	GetOrdersCountResponse struct {
		TotalCount int `json:"totalCount"`
	}
	GetOrdersByCustomerIdResponse struct {
		Items      []Order `json:"items"`
		TotalCount int     `json:"totalCount"`
	}
)
