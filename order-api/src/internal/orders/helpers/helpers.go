package helpers

import (
	"OrderAPI/src/internal/orders/handlers/types"
	"OrderAPI/src/pkg/errors"
	"fmt"
	"github.com/google/uuid"
)

func IsValidStatus(status string) bool {
	// Status Validation
	validOrderStatus := map[string]bool{
		"Accepted":  true,
		"Pending":   true,
		"Shipped":   true,
		"Delivered": true,
		"Canceled":  true,
	}
	if _, ok := validOrderStatus[status]; !ok {
		return false
	}
	return true
}

func ValidateChangeOrderStatusRequest(req types.ChangeOrderStatusRequest) error {
	var err error
	// OrderId validation
	_, err = uuid.Parse(req.OrderId)
	if err != nil {
		return errors.OrderUUIDInvalid
	}

	// Status Validation
	validOrderStatus := map[string]bool{
		"Accepted":  true,
		"Pending":   true,
		"Shipped":   true,
		"Delivered": true,
		"Canceled":  true,
	}
	if _, ok := validOrderStatus[req.Status]; !ok {
		return errors.ValidatorError.WrapDesc(fmt.Sprintf("Invalid Status value:%s. Valid Status values are Accepted,Pending,Shipped,Delivered,Canceled.", req.Status))
	}
	return nil
}
