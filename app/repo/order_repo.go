package repo

import (
	"webhook/domain"
	"webhook/pb"
)

type orderRepo struct {
	order pb.OrderServiceClient
}

func NewOrderRepo(order pb.OrderServiceClient) domain.OrderRepo {
	return &orderRepo{order: order}
}
