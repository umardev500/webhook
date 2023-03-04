package service

import (
	"context"
	"errors"
	"webhook/domain"
	"webhook/helper"
	"webhook/pb"
	"webhook/variable"
)

func (o *orderService) SetStatus(ctx context.Context, payload domain.SetStatusPayload) (ok bool, err error) {
	orderId := payload.OrderID
	status := payload.TransactionStatus

	if status == variable.PaymentStatusSettlement {
		settlementTimeStr := payload.SettlementTime
		settlementTime := helper.ParseLocalTime(settlementTimeStr)
		req := &pb.OrderChangeStatus{
			OrderId:        orderId,
			Status:         status,
			SettlementTime: settlementTime,
		}
		ok, err = o.repo.SetStatus(ctx, req)
		return
	}

	// will catch exire and cancel
	if status != variable.PaymentStatusPending {
		req := &pb.OrderChangeStatus{
			OrderId: orderId,
			Status:  status,
		}
		ok, err = o.repo.SetStatus(ctx, req)
		return
	}

	err = errors.New("not matching")
	return
}
