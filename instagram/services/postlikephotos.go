package services

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/instagram/models"
)


func CreatePost(user *models.User, url string) {
	photo := &models.Photo{
		URL: url,
		Owner: user,
	}
	post := &models.Post{
		Photo: photo,
		LikedBy: []*models.User{},
	}

	user.Post = append(user.Post, post)
}

func LikePost(user *models.User, post *models.Post) {
	for _, u := range post.LikedBy {
		if u==user {
			fmt.Println("Post Already Liked by user" + user.Name)
			return
		}
	}
	post.LikedBy = append(post.LikedBy, user)

	fmt.Println(user.Name + "Liked the post by" + post.Photo.Owner.Name)
}

func FollowUser(user *models.User, userToFollow *models.User) {
	for _, u := range user.Following {
		if u==userToFollow {
			fmt.Println(user.Name + "User is already Following the user" + userToFollow.Name)
			return
		}
	}
	user.Following = append(user.Following, userToFollow)
	
	fmt.Println(user.Name + "User started Following the user" + userToFollow.Name)
}

