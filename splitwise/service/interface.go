package service

import "github.com/rupeshkumarpallai/lld_go/splitwise/model"

type ExpenseServiceInterface interface {
	CreateExpense(expenseID string, paidBy *model.User, amount float64, participants []*model.User, date string) *model.Expense
	SplitExpense(expense *model.Expense) map[*model.User]float64
	AddExpenseToGroup(group *model.Group, expense *model.Expense)
}

type GroupServiceInterface interface {
	CreateGroup(groupID, groupName string, members []*model.User) *model.Group
	AddMember(group *model.Group, user *model.User)
	RemoveMember(group *model.Group, user *model.User)
	AddExpense(group *model.Group, expense *model.Expense)
}

type DebtServiceInterface interface {
	CreateDebt(owedBy *model.User, owedTo *model.User, amount float64) *model.Debt
	SettleDebt(debt *model.Debt, amount float64)
}

type PaymentServiceInterface interface {
	CreatePayment(paidBy *model.User, paidTo *model.User, amount float64) *model.Payment
	RecordPayment(payment *model.Payment, debt *model.Debt)
}

type UserServiceInterface interface {
	CreateUser(userID, name, email string) *model.User
	UpdateUser(user *model.User, name, email string)
}
