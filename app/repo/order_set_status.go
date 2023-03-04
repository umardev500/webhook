package repo

import (
	"context"
	"webhook/pb"
)

func (o *orderRepo) SetStatus(ctx context.Context, req *pb.OrderChangeStatus) (ok bool, err error) {
	resp, err := o.order.ChangeStatus(ctx, req)
	ok = resp.IsAffected
	return
}
