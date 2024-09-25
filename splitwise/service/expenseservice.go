package service

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/splitwise/model"
)

type ExpenseService struct{}

func (ex *ExpenseService) CreateExpense(expenseID string, paidBy *model.User, amount float64, participants []*model.User, date string) *model.Expense {
	expense := model.NewExpense(expenseID, paidBy, amount, participants, date)
	fmt.Printf("Expense created: %s, Paid by: %s, Amount: %.2f\n", expense.ExpenseID, paidBy.Name, amount)
	return expense
}

// Implementing the SplitExpense method
func (es *ExpenseService) SplitExpense(expense *model.Expense) map[*model.User]float64 {
	splitAmount := expense.SplitAmount() // Get even split for participants
	debtMap := make(map[*model.User]float64)

	for _, participant := range expense.Participants {
		if participant.UserID != expense.PaidBy.UserID {
			debtMap[participant] = splitAmount
		}
	}

	fmt.Printf("Expense split: %s, Split Amount: %.2f per person\n", expense.ExpenseID, splitAmount)
	return debtMap
}

func (es *ExpenseService) AddExpenseToGroup(group *model.Group, expense *model.Expense) {
	group.AddExpense(expense)
	fmt.Printf("Expense %s added to Group %s\n", expense.ExpenseID, group.GroupName)
}

