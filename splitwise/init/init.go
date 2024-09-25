package initializer

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/splitwise/model"
	"github.com/rupeshkumarpallai/lld_go/splitwise/service"
	"github.com/rupeshkumarpallai/lld_go/splitwise/utils"
)

// Initialize sets up initial data for the system
func Initialize() {
	// Initialize users
	user1 := model.NewUser("u1", "Alice", "alice@example.com")
	user2 := model.NewUser("u2", "Bob", "bob@example.com")
	user3 := model.NewUser("u3", "Charlie", "charlie@example.com")

	// Adding users to store (this simulates a database)
	utils.UsersStore[user1.UserID] = user1
	utils.UsersStore[user2.UserID] = user2
	utils.UsersStore[user3.UserID] = user3

	// Initialize group
	group := model.NewGroup("g1", "Friends", []*model.User{user1, user2, user3})
	
	// Add group and users to data store
	utils.GroupsStore[group.GroupID] = group

	// Initialize an ExpenseService
	expenseService := service.ExpenseService{}

	// Create expenses
	expense1 := expenseService.CreateExpense("e1", user1, 100.0, []*model.User{user1, user2, user3}, "2024-09-24")
	utils.ExpensesStore["e1"] = expense1
	expenseService.AddExpenseToGroup(group, expense1)

	fmt.Println("Initialization completed successfully!")
}
