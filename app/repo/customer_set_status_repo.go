package repo

import (
	"context"
	"webhook/pb"
)

func (c *customerRepo) SetStatus(ctx context.Context, req *pb.CustomerSetExpRequest) (res *pb.OperationResponse, err error) {
	res, err = c.customer.SetExp(ctx, req)
	return
}
