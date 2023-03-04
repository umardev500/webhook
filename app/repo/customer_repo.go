package repo

import (
	"webhook/domain"
	"webhook/pb"
)

type customerRepo struct {
	customer pb.CustomerServiceClient
}

func NewCustomerRepo(customer pb.CustomerServiceClient) domain.CustomerRepo {
	return &customerRepo{customer: customer}
}
