package utils

import (
	"CustomerAPI/src/internal/types"
	"CustomerAPI/src/pkg/errors"
	"github.com/google/uuid"
	"net/mail"
)

func ValidateCustomerInfos(c types.CustomerRequest) *errors.Error {

	//Name Validation
	if len(c.Name) <= 3 {
		return errors.CustomerNameInvalid
	}
	//Email Validation
	_, err := mail.ParseAddress(c.Email)
	if err != nil {
		return errors.CustomerEmailInvalid
	}
	//Address Validation
	if c.Address.AddressLine == "" || c.Address.City == "" || c.Address.Country == "" || c.Address.CityCode <= 0 {
		return errors.CustomerAddressInvalid
	}
	return nil
}

func ValidateCustomerInfosForUpdate(customer types.Customer) *errors.Error {

	//Check UUID Validation
	_, e := uuid.Parse(customer.Id)
	if e != nil {
		return errors.CustomerUUIDInvalid
	}

	if customer.Name != "" {
		//Name Validation
		if len(customer.Name) <= 3 {
			return errors.CustomerNameInvalid
		}
	}

	if customer.Email != "" {
		//Email Validation
		_, err := mail.ParseAddress(customer.Email)
		if err != nil {
			return errors.CustomerEmailInvalid
		}
	}

	return nil
}
