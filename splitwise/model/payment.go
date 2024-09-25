package model

type Payment struct {
	PaidBy *User
	PaidTo *User
	Amount float64
}

func NewPayment(paidBy *User, paidTo *User, amount float64) *Payment {
	return &Payment{
		PaidBy: paidBy,
		PaidTo: paidTo,
		Amount: amount,
	}
}


