package controller

import (
	"net/http"
	"webhook/domain"
	"webhook/variable"

	"github.com/gofiber/fiber/v2"
)

func (o *orderController) SetStatus(ctx *fiber.Ctx) (err error) {
	var payload domain.SetStatusPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	appCtx := ctx.Context()
	orderId := payload.OrderID
	status := payload.TransactionStatus
	_, err = o.service.SetStatus(appCtx, payload)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	if status == variable.PaymentStatusSettlement {
		// find order
		orderResp, err := o.service.GetOrder(appCtx, orderId)
		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}

		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}

		if orderResp.IsEmpty {
			return ctx.SendStatus(http.StatusNotFound)
		}

		order := orderResp.Payload
		userId := order.Buyer.CustomerId
		settlementTime := order.SettlementTime
		duration := order.Product.Duration

		customerStatusPayload := domain.CustomerStatusPayload{
			UserId:         userId,
			SettlementTime: settlementTime,
			Status:         "active",
			Duration:       int(duration),
		}
		_, err = o.customerService.SetStatus(appCtx, customerStatusPayload)
		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}
	}

	return ctx.SendStatus(http.StatusOK)
}
