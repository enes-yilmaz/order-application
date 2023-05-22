package helpers

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
