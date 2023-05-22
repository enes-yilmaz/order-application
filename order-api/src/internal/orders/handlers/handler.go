package handlers

import (
	"OrderAPI/src/internal/orders/clients"
	orderErrors "OrderAPI/src/internal/orders/errors"
	"OrderAPI/src/internal/orders/handlers/types"
	"OrderAPI/src/internal/orders/helpers"
	orderRepository "OrderAPI/src/internal/orders/storages/mongo"
	"OrderAPI/src/pkg/errors"
	"OrderAPI/src/pkg/utils"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	customerClient clients.Customer
	orderRepo      orderRepository.Repository
}

func NewHandler(g *echo.Group, orderRepository *orderRepository.Repository) *Handler {
	customerClient := clients.NewCustomerClient()
	orderRepo := orderRepository

	h := &Handler{customerClient: *customerClient, orderRepo: *orderRepo}

	g = g.Group("/orders")

	g.GET("", h.GetAllOrders)
	g.GET("/order/:orderId", h.GetOrderByOrderId)
	g.GET("/:customerId", h.GetOrdersByCustomerId)
	g.POST("/create", h.CreateOrder)
	g.POST("/update", h.UpdateOrder)
	g.DELETE("/:orderId", h.DeleteOrder)
	g.PATCH("/:orderId", h.ChangeOrderStatus)

	return h
}

// CreateOrder
// @Summary  Create an order
// @Tags     Orders
// @Accept   json
// @Produce  json
// @Param    RequestBody  body   types.CreateOrderRequest  true   " "
// @Success  201          "Created"
// @Failure  400          "Bad Request"
// @Failure  500          "Internal Error"
// @Router   /orders/create [post]
func (h Handler) CreateOrder(c echo.Context) error {
	req := new(types.CreateOrderRequest)
	var err error

	if _, err = utils.ValidateRequest(c, "create", req); err != nil {
		panic(err)
	}

	isValidCustomer, err := h.customerClient.ValidateCustomer(req.CustomerId)
	if err != nil {
		panic(err)
	}

	if isValidCustomer {
		order := types.Order{
			Id:         uuid.NewString(),
			CustomerId: req.CustomerId,
			Quantity:   req.Quantity,
			Price:      req.Price,
			Status:     req.Status,
			Address:    req.Address,
			Product:    req.Product,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		err = h.orderRepo.SaveOrder(order)
		if err != nil {
			return err.(*errors.Error).ToResponse(c)
		}
		return c.JSON(http.StatusCreated, order.Id)

	} else {
		return c.JSON(http.StatusBadRequest, orderErrors.NonExistCustomer)
	}

}

// UpdateOrder
// @Summary  Update an order
// @Tags     Orders
// @Accept   json
// @Produce  json
// @Param    RequestBody  body   types.UpdateOrderRequest  true   " "
// @Success  200          ""
// @Failure  400          "Bad Request"
// @Failure  500          "Internal Error"
// @Router   /orders/update [post]
func (h Handler) UpdateOrder(c echo.Context) error {
	req := new(types.UpdateOrderRequest)

	var err error

	if _, err = utils.ValidateRequest(c, "update", req); err != nil {
		panic(err)
	}

	isValidCustomer, err := h.customerClient.ValidateCustomer(req.CustomerId)
	if err != nil {
		panic(err)
	}

	if isValidCustomer {
		order := types.Order{
			Id:         req.OrderId,
			CustomerId: req.CustomerId,
			Quantity:   req.Quantity,
			Price:      req.Price,
			Status:     req.Status,
			Address:    req.Address,
			Product:    req.Product,
			UpdatedAt:  time.Now(),
		}
		modifiedCount, err := h.orderRepo.UpdateOrder(order)
		if err != nil {
			return err.(*errors.Error).ToResponse(c)
		}
		return c.JSON(http.StatusOK, modifiedCount > 0)

	} else {
		return c.JSON(http.StatusBadRequest, orderErrors.NonExistCustomer)
	}
}

// DeleteOrder
// @Summary  Delete an order
// @Tags     Orders
// @Accept   json
// @Produce  json
// @Param    orderId path	string true   " "
// @Success  204
// @Failure  400 "Bad Request"
// @Failure  500 "Internal Error"
// @Router   /orders/{orderId} [delete]
func (h Handler) DeleteOrder(c echo.Context) error {
	var err error

	orderId := strings.TrimSpace(c.Param("orderId"))
	_, err = uuid.Parse(orderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.OrderUUIDInvalid)
	}

	_, err = h.orderRepo.DeleteOrderByOrderId(orderId)
	if err != nil {
		panic(err)
	}
	return c.NoContent(http.StatusNoContent)
}

// ChangeOrderStatus
// @Summary  Update an order status
// @Tags     Orders
// @Accept   json
// @Produce  json
// @Param    orderId path	string true   " "
// @Param    RequestBody  body   types.ChangeOrderStatusRequest  true   " "
// @Success  204
// @Failure  400 "Bad Request"
// @Failure  500 "Internal Error"
// @Router   /orders/{orderId} [patch]
func (h Handler) ChangeOrderStatus(c echo.Context) error {
	var err error
	req := new(types.ChangeOrderStatusRequest)

	req.OrderId = strings.TrimSpace(c.Param("orderId"))
	_, err = uuid.Parse(req.OrderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.OrderUUIDInvalid)
	}
	if !helpers.IsValidStatus(req.Status) {
		return c.JSON(http.StatusBadRequest, errors.ValidatorError.WrapDesc(fmt.Sprintf("Invalid Status value:%s. Valid Status values are Accepted,Pending,Shipped,Delivered,Canceled.", req.Status)))

	}

	modifiedCount, err := h.orderRepo.ChangeOrderStatus(req)
	if err != nil {
		return err.(*errors.Error).ToResponse(c)
	}
	return c.JSON(http.StatusOK, modifiedCount > 0)
}

// GetAllOrders
// @Summary  Get all order
// @Tags     Orders
// @Accept   json
// @Produce  json
// @Param    limit     			query    number    	false " " "10"
// @Param    offset      		query    number		false " " "0"
// @Param    orderDirection  	query    string     false "orderDirection"
// @Param    orderBy         	query    string     false "orderBy"
// @Param    isCount         	query    boolean    false "false"
// @Success  200				""
// @Failure  404 				"Not Found"
// @Failure  500            	"Internal Error"
// @Router   /orders [get]
func (h Handler) GetAllOrders(c echo.Context) error {
	req := types.GetAllOrdersRequest{}

	if limit, offset, err := utils.ValidateLimitOffset(strings.TrimSpace(c.QueryParam("limit")), strings.TrimSpace(c.QueryParam("offset")), 10); err != nil {
		panic(err)
	} else {
		req.Limit, req.Offset = limit, offset
	}
	req.OrderDirection = strings.TrimSpace(c.QueryParam("orderDirection"))
	req.OrderBy = strings.TrimSpace(c.QueryParam("orderBy"))
	req.IsCount, _ = strconv.ParseBool(strings.TrimSpace(c.QueryParam("isCount")))

	orders, orderCount, err := h.orderRepo.GetAllOrders(req)
	if err != nil {
		panic(err)
	}
	if req.IsCount {
		return c.JSON(http.StatusOK, types.GetOrdersCountResponse{TotalCount: orderCount})
	}

	return c.JSON(http.StatusOK, types.GetOrdersResponse{Items: orders, TotalCount: orderCount, Limit: req.Limit, Offset: req.Offset})
}

// GetOrderByOrderId
// @Summary Get order by orderId
// @Tags    Orders
// @Accept  json
// @Produce json
// @Param   orderId		 path	string true   " "
// @Success 200			{object} types.Order
// @Failure 404 		"Not Found"
// @Failure 500 		"Internal Error"
// @Router  /orders/order/{orderId} [get]
func (h Handler) GetOrderByOrderId(c echo.Context) error {
	var err error
	var req types.GetOrderByOrderIdRequest

	req.OrderId = strings.TrimSpace(c.Param("orderId"))
	_, err = uuid.Parse(req.OrderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.OrderUUIDInvalid)
	}

	order, err := h.orderRepo.GetOrderByOrderId(req)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, order)
}

// GetOrdersByCustomerId
// @Summary Get order by customerId
// @Tags    Orders
// @Accept  json
// @Produce json
// @Param   customerId		 path	string true   " "
// @Success 200			{object} types.GetOrdersResponse{}
// @Failure 404 		"Not Found"
// @Failure 500 		"Internal Error"
// @Router  /orders/{customerId} [get]
func (h Handler) GetOrdersByCustomerId(c echo.Context) error {
	var err error
	var req types.GetOrdersByCustomerIdRequest

	req.CustomerId = strings.TrimSpace(c.Param("customerId"))
	_, err = uuid.Parse(req.CustomerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.OrderUUIDInvalid)
	}

	orders, err := h.orderRepo.GetOrdersByCustomerId(req)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, types.GetOrdersByCustomerIdResponse{Items: orders, TotalCount: len(orders)})
}
