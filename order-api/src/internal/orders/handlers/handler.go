package handlers

import (
	"OrderAPI/src/internal/orders/clients"
	orderErrors "OrderAPI/src/internal/orders/errors"
	"OrderAPI/src/internal/orders/handlers/types"
	orderRepository "OrderAPI/src/internal/orders/storages/mongo"
	"OrderAPI/src/pkg/errors"
	"OrderAPI/src/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
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

	g.POST("", h.CreateOrder)

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
// @Router   /orders [post]
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
