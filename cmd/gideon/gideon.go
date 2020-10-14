package main

import (
	"commander-list/controller"
	"commander-list/driver"
	"commander-list/model"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

var decks []model.Deck
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("hduahdueahduaeh")
	db = driver.ConnectDB()
	controller := controller.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/decks", controller.GetDecks(db)).Methods("GET")
	router.HandleFunc("/decks/{id}", controller.GetDeck(db)).Methods("GET")
	router.HandleFunc("/decks", controller.AddDeck(db)).Methods("POST")
	router.HandleFunc("/decks", controller.UpdateDeck(db)).Methods("PUT")
	router.HandleFunc("/decks/{id}", controller.RemoveDeck(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
