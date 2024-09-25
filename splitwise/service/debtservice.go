package service

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/splitwise/model"
)

type DebtService struct{}

// Method to create a new debt
func (ds *DebtService) CreateDebt(owedBy *model.User, owedTo *model.User, amount float64) *model.Debt {
	debt := model.NewDebt(owedBy, owedTo, amount)
	fmt.Printf("Debt created: %s owes %s %.2f\n", owedBy.Name, owedTo.Name, amount)
	return debt
}

// Method to settle a debt
func (ds *DebtService) SettleDebt(debt *model.Debt, amount float64) {
	debt.Settle(amount)
	fmt.Printf("Debt settled: %s paid %s %.2f\n", debt.OwedBy.Name, debt.OwedTo.Name, amount)
}