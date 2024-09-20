package services

import "github.com/rupeshkumarpallai/lld_go/instagram/models"

func GetFeed(user *models.User) []*models.Post {
	var feed []*models.Post

	for _, followedUsers := range user.Following {
		feed = append(feed, followedUsers.Post...)
	}

	return feed
}