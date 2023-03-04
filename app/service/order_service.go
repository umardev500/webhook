package service

import "webhook/domain"

type orderService struct {
	repo domain.OrderRepo
}

func NewOrderService(repo domain.OrderRepo) domain.OrderService {
	return &orderService{repo: repo}
}
