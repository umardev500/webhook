package service

import (
	"context"
	"webhook/domain"
	"webhook/pb"
)

func (c *customerService) SetStatus(ctx context.Context, payload domain.CustomerStatusPayload) (res *pb.OperationResponse, err error) {
	seconds := payload.Duration * 24 * 60 * 60
	estimatedExp := payload.SettlementTime + int64(seconds)

	req := &pb.CustomerSetExpRequest{
		CustomerId:     payload.UserId,
		Status:         payload.Status,
		SettlementTime: payload.SettlementTime,
		ExpTime:        int64(estimatedExp),
	}

	res, err = c.repo.SetStatus(ctx, req)

	return
}
