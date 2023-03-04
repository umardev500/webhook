package service

import "webhook/domain"

type customerService struct {
	repo domain.CustomerRepo
}

func NewCustomerService(repo domain.CustomerRepo) domain.CustomerService {
	return &customerService{repo: repo}
}
