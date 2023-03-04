package injector

import (
	"webhook/app/controller"
	"webhook/app/repo"
	"webhook/app/service"
	"webhook/pb"

	"github.com/gofiber/fiber/v2"
)

func NewOrderInjector(router fiber.Router, order pb.OrderServiceClient, customer pb.CustomerServiceClient) {
	orderRepo := repo.NewOrderRepo(order)
	orderService := service.NewOrderService(orderRepo)
	customreRepo := repo.NewCustomerRepo(customer)
	customerService := service.NewCustomerService(customreRepo)
	controller.NewOrderController(router, orderService, customerService)
}
