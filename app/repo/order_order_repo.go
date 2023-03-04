package repo

import (
	"context"
	"webhook/pb"
)

func (o *orderRepo) GetOrder(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error) {
	res, err = o.order.FindOne(ctx, req)
	return
}
