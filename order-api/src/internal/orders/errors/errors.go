package errors

import (
	"OrderAPI/src/pkg/errors"
	"net/http"
)

const repoOp = "orders.repo"
const handlerOp = "orders.handler"
const customerClientOp = "orders.client.customer"

var (
	AggregateError        = errors.New(repoOp, "AggregateError  ", 1001, http.StatusInternalServerError)
	CursorAllError        = errors.New(repoOp, "CursorAllError  ", 1002, http.StatusInternalServerError)
	InvalidCustomerId     = errors.New(handlerOp, "CustomerId Must Be UUID or can not empty ", 1003, http.StatusBadRequest)
	EmptyStartDate        = errors.New(handlerOp, "StartDate is can not empty", 1004, http.StatusBadRequest)
	EmptyEndDate          = errors.New(handlerOp, "EndDate is can not empty", 1005, http.StatusBadRequest)
	DateValidation        = errors.New(handlerOp, "Date validation error ", 1006, http.StatusInternalServerError)
	FailedToBindError     = errors.New(customerClientOp, "Failed to bind data from client", 10007, http.StatusInternalServerError)
	ValidateCustomerError = errors.New(customerClientOp, "ValidateCustomerError", 10008, http.StatusInternalServerError)
	NonExistCustomer      = errors.New(customerClientOp, "There is no customer with this customerId", 1009, http.StatusBadRequest)
	JsonUnmarshalFailed   = errors.New(handlerOp, "Json unmarshal error", 100010, http.StatusInternalServerError)
)
