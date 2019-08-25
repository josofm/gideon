package main

import (
	"commander-list/controller"
	"commander-list/driver"
	"commander-list/model"
	"database/sql"
	"encoding/json"
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
	db = driver.ConnectDB()
	controller := controller.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/decks", controller.GetDecks(db)).Methods("GET")
	router.HandleFunc("/decks/{id}", controller.GetDeck(db)).Methods("GET")
	router.HandleFunc("/decks", addDeck).Methods("POST")
	router.HandleFunc("/decks", updateDeck).Methods("PUT")
	router.HandleFunc("/decks/{id}", removeDeck).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func addDeck(w http.ResponseWriter, r *http.Request) {
	var deck model.Deck
	var deckID int

	json.NewDecoder(r.Body).Decode(&deck)

	err := db.QueryRow("insert into decks (owner, commander) values($1, $2) RETURNING id;",
		deck.Owner, deck.Commander).Scan(&deckID)
	logFatal(err)

	json.NewEncoder(w).Encode(deckID)

	log.Println("Add deck is called")
}

func updateDeck(w http.ResponseWriter, r *http.Request) {
	var deck model.Deck
	json.NewDecoder(r.Body).Decode(&deck)

	result, err := db.Exec("update decks set owner=$1, commander=$2 where id=$3 RETURNING id;",
		&deck.Owner, &deck.Commander, &deck.ID)
	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)

	log.Println("Update deck is called")
}

func removeDeck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := db.Exec("delete from decks where id = $1", params["id"])
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)

	log.Println("Delete deck is called")
}
