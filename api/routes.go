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

	//auth rote
	s := router.PathPrefix("/auth").Subrouter()
	s.Use(api.jwtVerify)
	s.HandleFunc("/user/{id}", api.getUser).Methods("GET")
	s.HandleFunc("/deck", api.addDeck).Methods("POST")
	s.HandleFunc("/user/{id}", api.deleteUser).Methods("DELETE")
	s.HandleFunc("/user/{id}", api.updateUser).Methods("PUT")

	//TO DO add admin route

	return router
}
