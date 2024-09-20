package models 

type User struct {
	Name string
	Following []*User
	Post []*Post
}

func NewUser(name string) *User {
	return &User{
		Name: name,
		Following: []*User{},
		Post: []*Post{},
	}
}