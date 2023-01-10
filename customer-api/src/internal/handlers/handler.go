package handlers

import (
	customerRepo "CustomerAPI/src/internal/storages/mongo"
	"CustomerAPI/src/internal/types"
	"CustomerAPI/src/pkg/errors"
	"CustomerAPI/src/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	customerRepo customerRepo.Repository
}

func NewHandler(baseGroup *echo.Group, customerRepo *customerRepo.Repository) *Handler {

	h := &Handler{customerRepo: *customerRepo}
	g := baseGroup.Group("/customers")

	g.GET("", h.GetAllCustomers)
	g.GET("/:customerId", h.GetCustomer)
	g.GET("/validate/:customerId", h.ValidateCustomer)
	g.POST("", h.CreateCustomer)
	g.POST("/update", h.UpdateCustomer)
	g.DELETE("/:customerId", h.DeleteCustomer)

	return h
}

// GetAllCustomers
// @Summary Get all customers
// @Tags    Customers
// @Accept  json
// @Produce json
// @Param   limit     		query    number    	false " " "10"
// @Param   offset      	query    number		false " " "0"
// @Param   orderDirection  query    string     false "orderDirection"
// @Param   orderBy         query    string     false "orderBy"
// @Param   isCount         query    boolean    false "false"
// @Success 200
// @Failure 404 "Not Found"
// @Failure 500 "Internal Error"
// @Router  /customers [get]
func (h Handler) GetAllCustomers(c echo.Context) error {
	var req types.GetAllCustomersRequest

	if limit, offset, err := utils.ValidateLimitOffset(strings.TrimSpace(c.QueryParam("limit")), strings.TrimSpace(c.QueryParam("offset")), 10); err != nil {
		panic(err)
	} else {
		req.Limit, req.Offset = limit, offset
	}
	req.OrderDirection = strings.TrimSpace(c.QueryParam("orderDirection"))
	req.OrderBy = strings.TrimSpace(c.QueryParam("orderBy"))
	req.IsCount, _ = strconv.ParseBool(strings.TrimSpace(c.QueryParam("isCount")))

	customers, err := h.customerRepo.GetAllCustomers(req)
	if err != nil {
		panic(err)
	}

	if req.IsCount {
		return c.JSON(http.StatusOK, types.GetCustomersCountResponse{TotalCount: customers.TotalCount[0].Count})
	}

	return c.JSON(http.StatusOK, types.GetCustomersResponse{Items: customers.Items, TotalCount: customers.TotalCount[0].Count, Limit: req.Limit, Offset: req.Offset})
}

// GetCustomer
// @Summary Get customer by customerId
// @Tags    Customers
// @Accept  json
// @Produce json
// @Param   customerId path	string true   " "
// @Success 200	{object} types.Customer
// @Failure 404 "Not Found"
// @Failure 500 "Internal Error"
// @Router  /customers/{customerId} [get]
func (h Handler) GetCustomer(c echo.Context) error {
	var err error
	var req types.GetCustomerRequest

	req.CustomerId = strings.TrimSpace(c.Param("customerId"))
	_, err = uuid.Parse(req.CustomerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.CustomerUUIDInvalid)
	}

	customer, err := h.customerRepo.GetCustomer(req)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, customer)
}

// GetCustomerValidation
// @Summary Get customer validation status
// @Tags    Customers
// @Accept  json
// @Produce json
// @Param   customerId path	string true   " "
// @Success 200
// @Failure 404 "Not Found"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Error"
// @Router  /customers/validate/{customerId} [get]
func (h Handler) ValidateCustomer(c echo.Context) error {
	var err error
	var req types.GetCustomerRequest

	req.CustomerId = strings.TrimSpace(c.Param("customerId"))
	_, err = uuid.Parse(req.CustomerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.CustomerUUIDInvalid)
	}

	isValid, err := h.customerRepo.IsValidCustomer(req)
	if err != nil {
		panic(err)
	}

	if isValid {
		return c.JSON(http.StatusOK, map[string]bool{
			"isValidCustomer": true,
		})
	} else {
		return c.JSON(http.StatusNotFound, map[string]bool{
			"isValidCustomer": false,
		})
	}

}

// CreateCustomer
// @Summary  Create a customer
// @Tags     Customers
// @Accept   json
// @Produce  json
// @Param    RequestBody  body   types.CreateCustomerRequest  true   " "
// @Success  201          "Created"
// @Failure  400          "Bad Request"
// @Failure  500          "Internal Error"
// @Router   /customers [post]
func (h Handler) CreateCustomer(c echo.Context) error {
	req := new(types.CreateCustomerRequest)
	var err error
	if _, err = utils.ValidateRequest(c, "create", req); err != nil {
		panic(err)
	}
	if isExist := h.customerRepo.IsCustomerExist(req.Email); isExist {
		return c.JSON(http.StatusBadRequest, errors.CustomerAlreadyCreated)
	}

	customer := types.Customer{
		Id:        uuid.NewString(),
		Name:      req.Name,
		Email:     req.Email,
		Address:   req.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = h.customerRepo.SaveCustomer(customer)
	if err != nil {
		return err.(*errors.Error).ToResponse(c)
	}

	return c.JSON(http.StatusCreated, customer.Id)

}

// UpdateCustomer
// @Summary  Update a customer
// @Tags     Customers
// @Accept   json
// @Produce  json
// @Param    RequestBody  body   types.UpdateCustomerRequest  true   " "
// @Success  201          "Created"
// @Failure  400          "Bad Request"
// @Failure  500          "Internal Error"
// @Router   /customers/update [post]
func (h Handler) UpdateCustomer(c echo.Context) error {
	req := new(types.UpdateCustomerRequest)
	var err error
	if _, err = utils.ValidateRequest(c, "update", req); err != nil {
		panic(err)
	}

	filter := bson.M{"_id": strings.TrimSpace(req.Id)}

	// D struct , M key value , A array,
	updateModel := bson.M{
		"updatedAt": time.Now(),
	}
	if strings.TrimSpace(req.Name) != "" {
		updateModel["name"] = strings.TrimSpace(req.Name)
	}
	if strings.TrimSpace(req.Email) != "" {
		updateModel["email"] = strings.TrimSpace(req.Email)
	}

	if req.Address != (types.Address{}) {
		updateModel["address"] = req.Address
	}

	updateModel = bson.M{"$set": updateModel}

	//err = h.customerRepo.UpdateOneByCustomerId(filter, updateModel)
	customer, err := h.customerRepo.UpdateOneByCustomerId2(filter, updateModel)
	if err != nil {
		return err.(*errors.Error).ToResponse(c)
	}

	return c.JSON(http.StatusCreated, customer)
	//return c.NoContent(http.StatusNoContent)
}

// DeleteCustomer
// @Summary Delete customer by customerId
// @Tags    Customers
// @Accept  json
// @Produce json
// @Param   customerId path	string true   " "
// @Success 204
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Error"
// @Router  /customers/{customerId} [delete]
func (h Handler) DeleteCustomer(c echo.Context) error {
	var err error

	customerId := strings.TrimSpace(c.Param("customerId"))
	_, err = uuid.Parse(customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.CustomerUUIDInvalid)
	}

	_, err = h.customerRepo.DeleteCustomerById(customerId)
	if err != nil {
		panic(err)
	}

	return c.NoContent(http.StatusNoContent)
}
