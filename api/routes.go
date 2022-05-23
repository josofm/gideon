package api

import (
	"github.com/gorilla/mux"
)

func (api *Api) routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use()

	router.HandleFunc("/up", api.Up).Methods("GET")
	router.HandleFunc("/login", api.login).Methods("POST")
	router.HandleFunc("/register", api.register).Methods("POST")
	router.HandleFunc("/card/{name}", api.getCardByName).Methods("GET")

	router.HandleFunc("/user/{id}", api.jwtVerify(api.getUser)).Methods("GET")
	router.HandleFunc("/deck", api.jwtVerify(api.addDeck)).Methods("POST")
	router.HandleFunc("/user/{id}", api.jwtVerify(api.deleteUser)).Methods("DELETE")
	router.HandleFunc("/user/{id}", api.jwtVerify(api.updateUser)).Methods("PATCH")

	return router
}
