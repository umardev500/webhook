package controller

import (
	"webhook/domain"

	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	service         domain.OrderService
	customerService domain.CustomerService
}

func NewOrderController(router fiber.Router, service domain.OrderService, customerService domain.CustomerService) {
	handler := &orderController{
		service:         service,
		customerService: customerService,
	}
	router.Post("/status", handler.SetStatus)
}
