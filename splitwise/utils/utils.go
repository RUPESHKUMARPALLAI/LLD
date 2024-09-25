package utils

import "github.com/rupeshkumarpallai/lld_go/splitwise/model"

// Simulated user and expense data stores
var UsersStore = map[string]*model.User{}
var ExpensesStore = map[string]*model.Expense{}
var GroupsStore = map[string]*model.Group{}

// GetUserByID fetches a user by their ID
func GetUserByID(userID string) *model.User {
	return UsersStore[userID]
}

// GetUsersByIDs fetches multiple users by their IDs
func GetUsersByIDs(userIDs []string) []*model.User {
	users := []*model.User{}
	for _, userID := range userIDs {
		if user := GetUserByID(userID); user != nil {
			users = append(users, user)
		}
	}
	return users
}

// GetExpenseByID fetches an expense by its ID
func GetExpenseByID(expenseID string) *model.Expense {
	return ExpensesStore[expenseID]
}

// GetGroupByID fetches a group by its ID
func GetGroupByID(groupID string) *model.Group {
	return GroupsStore[groupID]
}