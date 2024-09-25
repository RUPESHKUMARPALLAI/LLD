package model

type Group struct {
	GroupID string
	GroupName string
	Members []*User
	Expenses []*Expense
}

func NewGroup(groupID, groupName string, members []*User) *Group {
	return &Group{
		GroupID: groupID,
		GroupName: groupName,
		Members: members,
	}
}

func (g *Group) AddMemeber(user *User) {
	g.Members = append(g.Members, user)
}

func (g *Group) RemoveMemeber(user *User) {
	for i, member := range g.Members {
		if member.UserID == user.UserID {
			g.Members = append(g.Members[:i],g.Members[i+1:]...)
			break;
		} 
	}
}

func (g *Group) AddExpense(expense *Expense) {
	g.Expenses = append(g.Expenses, expense)
}