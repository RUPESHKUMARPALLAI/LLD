package service

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/splitwise/model"
)

// PaymentService manages payment-related operations
type PaymentService struct{}

// Method to create a payment
func (ps *PaymentService) CreatePayment(paidBy *model.User, paidTo *model.User, amount float64) *model.Payment {
	payment := model.NewPayment(paidBy, paidTo, amount)
	fmt.Printf("Payment created: %s paid %s %.2f\n", paidBy.Name, paidTo.Name, amount)
	return payment
}

// Method to record payment and settle related debt
func (ps *PaymentService) RecordPayment(payment *model.Payment, debt *model.Debt) {
	debt.Settle(payment.Amount)
	fmt.Printf("Payment recorded: %s paid %s %.2f and settled debt\n", payment.PaidBy.Name, payment.PaidTo.Name, payment.Amount)
}
