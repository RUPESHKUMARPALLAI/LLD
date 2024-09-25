package model

type Expense struct {
	ExpenseID string
	PaidBy *User
	Amount float64
	Participants []*User
	Date string
}

func NewExpense(expenseID string, paidBy *User, amount float64, participants []*User, date string) *Expense {
	return &Expense{
		ExpenseID: expenseID,
		PaidBy: paidBy,
		Amount: amount,
		Participants: participants,
		Date: date,
	}
}

func (e *Expense) SplitAmount() float64 {
	if len(e.Participants) == 0 {
		return 0;
	}
	return e.Amount/float64(len(e.Participants))
}