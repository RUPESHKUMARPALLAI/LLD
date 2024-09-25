package model

type Debt struct {
	OwedBy *User
	OwedTo *User
	Amount float64
}

func NewDebt(owedBy *User, owedTo *User, amount float64) *Debt {
	return &Debt{
		OwedBy: owedBy,
		OwedTo: owedTo,
		Amount: amount,
	}
}

func (d *Debt) Settle(amount float64) {
	d.Amount -= amount
}