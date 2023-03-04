package service

import (
	"context"
	"webhook/pb"
)

func (o *orderService) GetOrder(ctx context.Context, orderId string) (res *pb.OrderFindOneResponse, err error) {
	res, err = o.repo.GetOrder(ctx, &pb.OrderFindOneRequest{OrderId: orderId})
	return
}
