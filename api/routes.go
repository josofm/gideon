package api

import (
	"github.com/gorilla/mux"
)

func (api *Api) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/up", api.Up).Methods("GET")
	router.HandleFunc("/login", api.login).Methods("POST")

	return router
}
