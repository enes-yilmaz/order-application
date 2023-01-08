package errors

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

const customerApi string = "customer-api"
const validatorOp string = "customer-api.handlers.validator"
const repoOp string = "customer-api.repository"

var (
	UnknownError                          = New(customerApi, "Unknown Error !!!", 10500, http.StatusInternalServerError)
	FailedImplementParamIntoTemplateError = New(customerApi, "Failed to implement given parameter into query", 10501, http.StatusInternalServerError)
)

var (
	ValidatorError = New(validatorOp, "", 10400, http.StatusBadRequest)
)

var (
	UpdateOneFailed   = New(repoOp, "Update operation failed", 10502, http.StatusInternalServerError)
	DeleteOneFailed   = New(repoOp, "Delete operation failed", 10505, http.StatusInternalServerError)
	FindFailed        = New(repoOp, "Finding operation failed", 10503, http.StatusInternalServerError)
	MongoCursorFailed = New(repoOp, "Mongo cursor Failed", 10504, http.StatusInternalServerError)
)

const customerServiceApiOp string = "customer-service.handler"

var (
	//CustomerInfoInvalid    = error.New(customerServiceApiOp, "Invalid Customer Infos are given", 1002, http.StatusBadRequest)
	CustomerNameInvalid    = New(customerServiceApiOp, "Invalid Customer name given, name length should be greater tha 3!", 1001, http.StatusBadRequest)
	CustomerEmailInvalid   = New(customerServiceApiOp, "Invalid Customer email given", 1002, http.StatusBadRequest)
	CustomerAddressInvalid = New(customerServiceApiOp, "Invalid Customer address given", 1003, http.StatusBadRequest)

	CustomerUUIDInvalid = New(customerServiceApiOp, "Invalid Customer uuid given", 1004, http.StatusBadRequest)
	CustomerNotFound    = New(customerServiceApiOp, "Customer Not Found", 1005, http.StatusNotFound)

	CustomerAlreadyCreated = New(customerServiceApiOp, "Customer with this email is exist.", 1015, http.StatusBadRequest)

	LimitInvalid  = New(customerServiceApiOp, "Invalid parameter - limit", 1006, http.StatusBadRequest)
	OffsetInvalid = New(customerServiceApiOp, "Invalid parameter - offset", 1007, http.StatusBadRequest)

	ModelParseError = New(customerServiceApiOp, "Model parse error", 1008, http.StatusInternalServerError)
)

// Pagination
var (
	LimitLessThanZeroError  = New(validatorOp, "Limit value must be greater than zero", 4001, http.StatusBadRequest)
	OffsetLessThanZeroError = New(validatorOp, "Offset value must be greater than zero", 4002, http.StatusBadRequest)
	LimitParsingError       = New(validatorOp, "Limit value couldn't parsed", 4002, http.StatusBadRequest)
	OffsetParsingError      = New(validatorOp, "Offset value couldn't parsed", 4002, http.StatusBadRequest)
)

// Time
var (
	EndDateTimeParsingError   = New(validatorOp, "Couldn't parse endDate to time package", 2000, 400)
	EndDateGreaterError       = New(validatorOp, "endDate can not be greater than today", 2001, 400)
	StartDateTimeParsingError = New(validatorOp, "Couldn't parse startDate to time package", 2003, 400)
	StartDateGreaterError     = New(validatorOp, "startDate can not be greater than endDate", 2004, 400)
	StartGreaterThanEndError  = New(validatorOp, "startDate can not be greater than endDate", 2005, 400)
)

type Error struct {
	Public     PublicError
	StatusCode int
	Internal   error
	Args       interface{}
}

type PublicError struct {
	Op        string
	Desc      string
	ErrorCode int
}

func (e *Error) Error() string {
	return fmt.Sprintf("Operation: %s, Description: %s, ErrorCode: %d, Internal: %v , Args: %v", e.Public.Op, e.Public.Desc, e.Public.ErrorCode, e.Internal, e.Args)
}

func New(op string, desc string, errorCode int, statusCode int) *Error {
	return &Error{Public: PublicError{
		Op:        op,
		Desc:      desc,
		ErrorCode: errorCode,
	}, StatusCode: statusCode}
}

func (e *Error) WrapDesc(desc string) *Error {
	return &Error{Public: PublicError{
		Op:        e.Public.Op,
		Desc:      desc,
		ErrorCode: e.Public.ErrorCode,
	},
		StatusCode: e.StatusCode,
	}
}

func (e *Error) Wrap(err error, args ...interface{}) *Error {
	if err == nil {
		return nil
	}

	return &Error{Public: PublicError{
		Op:        e.Public.Op,
		Desc:      e.Public.Desc,
		ErrorCode: e.Public.ErrorCode,
	},
		StatusCode: e.StatusCode,
		Internal:   err,
		Args:       args,
	}
}

func (e Error) ToResponse(c echo.Context) error {
	return c.JSON(e.StatusCode, e.Public)
}

func (e *Error) Log() {
	fields := logrus.Fields{
		"StatusCode": e.StatusCode,
		"Op":         e.Public.Op,
		"ErrorCode":  e.Public.ErrorCode,
		"Args":       e.Args,
		"Internal":   e.Internal,
	}
	if e.StatusCode >= 400 && e.StatusCode < 500 {
		logrus.WithFields(fields).Info(e.Public.Desc)
	} else if e.StatusCode >= 500 {
		logrus.WithFields(fields).Error(e.Public.Desc)
	}
}
