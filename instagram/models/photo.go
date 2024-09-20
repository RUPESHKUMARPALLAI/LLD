package models

type Photo struct {
	URL string
	Owner *User
}

type Post struct {
	Photo *Photo
	LikedBy []*User
}