package domain

import (
	"context"
	"webhook/pb"
)

type CustomerStatusPayload struct {
	UserId         string
	SettlementTime int64
	Status         string
	Duration       int
}

type CustomerService interface {
	SetStatus(ctx context.Context, payload CustomerStatusPayload) (res *pb.OperationResponse, err error)
}
type CustomerRepo interface {
	SetStatus(ctx context.Context, req *pb.CustomerSetExpRequest) (res *pb.OperationResponse, err error)
}
