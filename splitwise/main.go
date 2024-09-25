package main

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/splitwise/init"
	"github.com/rupeshkumarpallai/lld_go/splitwise/model"
	"github.com/rupeshkumarpallai/lld_go/splitwise/service"
	"github.com/rupeshkumarpallai/lld_go/splitwise/utils"
)

func main() {
	initializer.Initialize()

	// Example usage of services after initialization
	expenseService := service.ExpenseService{}

	// Fetch the group and list its expenses
	group := utils.GroupsStore["g1"]
	fmt.Printf("Expenses for group: %s\n", group.GroupName)
	for _, expense := range group.Expenses {
		fmt.Printf("Expense: %s, Paid by: %s, Amount: %.2f\n", expense.ExpenseID, expense.PaidBy.Name, expense.Amount)
	}

	// Fetch the first expense (simulated database retrieval)
	expense := utils.GetExpenseByID("e1")
	if expense == nil {
		panic("Expense with ID e1 not found!")
	}
	

	// Split the expense and log the debts
	debts := expenseService.SplitExpense(expense)
	for user, amount := range debts {
		fmt.Printf("User %s owes %.2f\n", user.Name, amount)
	}

	// Example to settle debt between Alice and Bob
	user1 := utils.GetUserByID("u1")
	user2 := utils.GetUserByID("u2")

	debt := model.NewDebt(user2, user1, 50) // Bob owes Alice $50
	fmt.Printf("%s owes %s %.2f\n", debt.OwedBy.Name, debt.OwedTo.Name, debt.Amount)

	// Settle the debt
	debt.Settle(50)
	fmt.Printf("After settlement, %s owes %s %.2f\n", debt.OwedBy.Name, debt.OwedTo.Name, debt.Amount)
}
