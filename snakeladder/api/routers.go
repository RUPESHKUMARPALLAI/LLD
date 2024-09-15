package api

import "github.com/gorilla/mux"

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/play", PlayHandler).Methods("GET")

	return r
}
