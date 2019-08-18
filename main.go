package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
	//"strconv"
)

type Deck struct {
	ID        int    `json:id`
	Commander string `json:commander` //will be type card
	Owner     string `json:owner`     //will be type user
}

var decks []Deck
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
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()

	router.HandleFunc("/decks", getDecks).Methods("GET")
	router.HandleFunc("/decks/{id}", getDeck).Methods("GET")
	router.HandleFunc("/decks", addDeck).Methods("POST")
	router.HandleFunc("/decks", updateDeck).Methods("PUT")
	router.HandleFunc("/decks/{id}", removeDeck).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getDecks(w http.ResponseWriter, r *http.Request) {
	var deck Deck
	decks = []Deck{}

	rows, err := db.Query("select * from decks")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&deck.ID, &deck.Commander, &deck.Owner)
		logFatal(err)

		decks = append(decks, deck)
	}

	json.NewEncoder(w).Encode(decks)

	log.Println("Get decks is called")
}

func getDeck(w http.ResponseWriter, r *http.Request) {
	var deck Deck
	params := mux.Vars(r)

	rows := db.QueryRow("select * from decks where id=$1", params["id"])
	err := rows.Scan(&deck.ID, &deck.Commander, &deck.Owner)
	logFatal(err)

	json.NewEncoder(w).Encode(deck)

	log.Println("Get deck is called")
}

func addDeck(w http.ResponseWriter, r *http.Request) {
	var deck Deck
	var deckID int

	json.NewDecoder(r.Body).Decode(&deck)

	err := db.QueryRow("insert into decks (owner, commander) values($1, $2) RETURNING id;",
		deck.Owner, deck.Commander).Scan(&deckID)
	logFatal(err)

	json.NewEncoder(w).Encode(deckID)

	log.Println("Add deck is called")
}

func updateDeck(w http.ResponseWriter, r *http.Request) {
	var deck Deck
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
