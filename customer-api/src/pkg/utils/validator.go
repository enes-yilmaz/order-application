package utils

import (
	"CustomerAPI/src/internal/types"
	"CustomerAPI/src/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/mail"
	"strconv"
	"strings"
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
		v.RegisterStructValidation(createCustomerValidations, types.CreateCustomerRequest{})
	} else if operation == "update" {
		v.RegisterStructValidation(updateCustomerValidations, types.UpdateCustomerRequest{})
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

func createCustomerValidations(sl validator.StructLevel) {
	model := sl.Current().Interface().(types.CreateCustomerRequest)

	if strings.TrimSpace(model.Name) == "" || strings.TrimSpace(model.Email) == "" {
		panic(errors.ValidatorError.WrapDesc("Name and email parameters are mandatory!"))
	}

	//Name Validation
	if len(model.Name) <= 3 {
		panic(errors.CustomerNameInvalid)
	}

	//Email Validation
	_, err := mail.ParseAddress(model.Email)
	if err != nil {
		panic(errors.CustomerEmailInvalid)
	}

}

func updateCustomerValidations(sl validator.StructLevel) {
	model := sl.Current().Interface().(types.UpdateCustomerRequest)

	//Check UUID Validation
	_, e := uuid.Parse(model.Id)
	if e != nil {
		panic(errors.CustomerUUIDInvalid)
	}

	//Name Validation
	if len(strings.TrimSpace(model.Name)) <= 3 {
		panic(errors.CustomerNameInvalid)
	}

	//Email Validation
	_, err := mail.ParseAddress(strings.TrimSpace(model.Email))
	if err != nil {
		panic(errors.CustomerEmailInvalid)
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
