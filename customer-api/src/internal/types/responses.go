package types

type (
	TotalCount struct {
		Count int `bson:"count" json:"count"`
	}
	Customers struct {
		Items      []Customer   `bson:"items" json:"items"`
		TotalCount []TotalCount `bson:"totalCount" json:"totalCount"`
	}
	GetCustomersResponse struct {
		Items      []Customer `json:"items"`
		TotalCount int        `json:"totalCount"`
		Limit      int        `json:"limit"`
		Offset     int        `json:"offset"`
	}
	GetCustomersCountResponse struct {
		TotalCount int `json:"totalCount"`
	}
)
