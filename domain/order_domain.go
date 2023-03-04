package domain

import (
	"context"
	"webhook/pb"
)

type SetStatusPayload struct {
	VaNumbers []struct {
		VaNumber string `json:"va_number,omitempty"`
		Bank     string `json:"bank,omitempty"`
	} `json:"va_numbers,omitempty"`
	TransactionTime   string `json:"transaction_time,omitempty"`
	TransactionStatus string `json:"transaction_status,omitempty"`
	TransactionID     string `json:"transaction_id,omitempty"`
	StatusMessage     string `json:"status_message,omitempty"`
	StatusCode        string `json:"status_code,omitempty"`
	SignatureKey      string `json:"signature_key,omitempty"`
	SettlementTime    string `json:"settlement_time,omitempty"`
	PermataVaNumber   string `json:"permata_va_number,omitempty"`
	PaymentType       string `json:"payment_type,omitempty"`
	PaymentAmounts    []struct {
		PaidAt string `json:"paid_at,omitempty"`
		Amount string `json:"amount,omitempty"`
	} `json:"payment_amounts,omitempty"`
	OrderID     string `json:"order_id,omitempty"`
	MerchantID  string `json:"merchant_id,omitempty"`
	GrossAmount string `json:"gross_amount,omitempty"`
	FraudStatus string `json:"fraud_status,omitempty"`
	Currency    string `json:"currency,omitempty"`
}

type StatusPayload struct {
	OrderId        string
	Status         string
	SettlementTime int64
}

type OrderService interface {
	SetStatus(ctx context.Context, payload SetStatusPayload) (bool, error)
	GetOrder(ctx context.Context, orderId string) (res *pb.OrderFindOneResponse, err error)
}

type OrderRepo interface {
	SetStatus(ctx context.Context, req *pb.OrderChangeStatus) (bool, error)
	GetOrder(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error)
}
