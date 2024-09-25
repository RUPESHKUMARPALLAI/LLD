package service

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/splitwise/model"
)

type GroupService struct{}

func (gs *GroupService) AddMember(group *model.Group, user *model.User) {
	group.AddMemeber(user)
	fmt.Printf("User %s added to Group %s\n", user.Name, group.GroupName)
}

func (gs *GroupService) RemoveMember(group *model.Group, user *model.User) {
	group.RemoveMemeber(user)
	fmt.Printf("User %s removed from Group %s\n", user.Name, group.GroupName)
}

func (gs *GroupService) AddExpense(group *model.Group, expense *model.Expense) {
	group.AddExpense(expense)
	fmt.Printf("Expense %s added to Group %s\n", expense.ExpenseID, group.GroupName)
}

func CreateGroup(groupID, groupName string, members []*model.User) *model.Group {
	return model.NewGroup(groupID, groupName, members)
}