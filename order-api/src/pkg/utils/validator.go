package utils

import (
	"OrderAPI/src/internal/orders/handlers/types"
	"OrderAPI/src/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func ValidateLimitOffset(limit, offset string, defaultLimit int) (int, int, error) {
	var lmt int
	var offst int

	if len(limit) > 0 {

		if l, err := strconv.Atoi(limit); l <= 0 && err == nil {
			return 0, 0, errors.LimitLessThanZeroError
		}

		l, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return 0, 0, errors.LimitParsingError
		}

		if l > 300 {
			l = 300
		}

		lmt = int(l)

	} else {
		lmt = defaultLimit
	}

	if len(offset) > 0 {
		o, err := strconv.ParseInt(offset, 10, 64)
		if err != nil {
			return 0, 0, errors.OffsetParsingError
		}
		if o < 0 {
			return 0, 0, errors.OffsetLessThanZeroError
		}
		offst = int(o)
	}

	return lmt, offst, nil
}

func ValidateRequest(ctx echo.Context, operation string, request interface{}) (result interface{}, err error) {

	err = ctx.Bind(request)
	if err != nil {
		if _, ok := err.(*echo.HTTPError); ok {
			return nil, errors.ValidatorError.WrapDesc(err.(*echo.HTTPError).Message.(string))
		}
		return errors.UnknownError.Wrap(err), nil
	}

	v := validator.New()

	if operation == "create" {
		v.RegisterStructValidation(createOrderValidations, types.CreateOrderRequest{})
	} else if operation == "update" {
		v.RegisterStructValidation(updateOrderValidations, types.UpdateOrderRequest{})
	}

	if err = v.Struct(request); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			desc := ""
			for _, err := range err.(validator.ValidationErrors) {
				desc = fmt.Sprintf("Problem is: '%s' , Hint: '%s' ", err.Field(), err.ActualTag())
			}
			return nil, errors.ValidatorError.WrapDesc(desc)
		}
		return errors.UnknownError.Wrap(err), nil
	}

	return request, nil
}

func createOrderValidations(sl validator.StructLevel) {
	model := sl.Current().Interface().(types.CreateOrderRequest)

	// Status Validation
	validOrderStatus := map[string]bool{
		"Accepted":  true,
		"Pending":   true,
		"Shipped":   true,
		"Delivered": true,
		"Canceled":  true,
	}
	if _, ok := validOrderStatus[model.Status]; !ok {
		panic(errors.ValidatorError.WrapDesc(fmt.Sprintf("Invalid Status value:%s. Valid Status values are Accepted,Pending,Shipped,Delivered,Canceled.", model.Status)))

	}

	//Price Validation
	if model.Price < 0 {
		panic(errors.ValidatorError.WrapDesc("Price cannot be less than 0!"))
	}

	//Quantity Validation
	if model.Quantity < 1 {
		panic(errors.ValidatorError.WrapDesc("Quantity should be greater than 0!"))
	}

}

func updateOrderValidations(sl validator.StructLevel) {
	model := sl.Current().Interface().(types.UpdateOrderRequest)

	// Status Validation
	validOrderStatus := map[string]bool{
		"Accepted":  true,
		"Pending":   true,
		"Shipped":   true,
		"Delivered": true,
		"Canceled":  true,
	}
	if _, ok := validOrderStatus[model.Status]; !ok {
		panic(errors.ValidatorError.WrapDesc(fmt.Sprintf("Invalid Status value:%s Valid Status values are Accepted,Pending,Shipped,Delivered,Canceled.", model.Status)))

	}

	//Price Validation
	if model.Price < 0 {
		panic(errors.ValidatorError.WrapDesc("Price cannot be less than 0!"))
	}

	//Quantity Validation
	if model.Quantity < 1 {
		panic(errors.ValidatorError.WrapDesc("Quantity should be greater than 0!"))
	}

}

func ValidateTime(startDateParam, endDateParam, format string, interval time.Duration) (time.Time, time.Time, error) {
	if startDateParam == "" {
		end := time.Now()
		startDateParam = end.Add(interval).Format(format)
	}
	if endDateParam == "" {
		now := time.Now()
		endDateParam = now.Format(format)

	}
	var err error

	start, err := time.Parse(format, startDateParam)
	if err != nil {
		return time.Time{}, time.Time{}, errors.StartDateTimeParsingError
	}

	end, err := time.Parse(format, endDateParam)
	if err != nil {
		return time.Time{}, time.Time{}, errors.EndDateTimeParsingError
	}

	now := time.Now().Format(format)
	formatNow, _ := time.Parse(format, now)

	if start.After(formatNow) {
		return time.Time{}, time.Time{}, errors.StartDateGreaterError
	}

	if end.After(formatNow) {
		return time.Time{}, time.Time{}, errors.EndDateGreaterError
	}

	if end.Before(start) {
		return time.Time{}, time.Time{}, errors.StartGreaterThanEndError
	}

	return start, end, nil
}
