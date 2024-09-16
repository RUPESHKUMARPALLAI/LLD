package main

import (
	"fmt"
	"net/http"

	"github.com/snakeladder/api"
	"github.com/snakeladder/engine"
)

func main() {
	gameEngine := engine.InitEngine(1, 1, 100)

	gameEngine.AddPlayer("Ram")
	gameEngine.AddPlayer("Shyam")

	api.SetGameEngine(gameEngine)

	//settingup router
	r := api.SetupRouter()

	//starting server
	port := "8080"
	fmt.Printf("starting server on port %s......\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}

}
