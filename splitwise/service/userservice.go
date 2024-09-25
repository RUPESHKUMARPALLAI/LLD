package service

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/splitwise/model"
)

// UserService manages user-related operations
type UserService struct{}

// Method to create a new user
func (us *UserService) CreateUser(userID, name, email string) *model.User {
	user := model.NewUser(userID, name, email)
	fmt.Printf("User created: %s, Email: %s\n", name, email)
	return user
}

// Method to update an existing user
func (us *UserService) UpdateUser(user *model.User, name, email string) {
	user.UpdateUser(name, email)
	fmt.Printf("User updated: %s, Email: %s\n", name, email)
}
