package main

import (
	"fmt"

	"github.com/rupeshkumarpallai/lld_go/instagram/models"
	"github.com/rupeshkumarpallai/lld_go/instagram/services"
)

func main() {
	
	user1 := models.NewUser("Ram")
	user2 := models.NewUser("Shyam")
	user3 := models.NewUser("Hari")

	services.CreatePost(user1, "jpg1")
	services.CreatePost(user1, "jpg2")
	services.CreatePost(user2, "jpg1")

	services.FollowUser(user2, user1)
	services.FollowUser(user3, user1)

	feed := services.GetFeed(user1)
	
	for _, post := range feed {
		fmt.Println(post.Photo.URL)
	}
}