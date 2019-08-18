package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Deck struct {
	ID        int    `json:id`
	Commander string `json:commander` //will be type card
	Owner     string `json:owner`     //will be type user
}

var decks []Deck

func main() {

	decks = append(decks,
		Deck{ID: 1, Commander: "Scarab-God", Owner: "Mirko"},
		Deck{ID: 2, Commander: "Borborigmo", Owner: "Joso"},
		Deck{ID: 3, Commander: "Saskia", Owner: "Jesus"},
		Deck{ID: 4, Commander: "Atraxa", Owner: "Graia"},
		Deck{ID: 5, Commander: "Muldrotha", Owner: "Jonatan"},
	)

	router := mux.NewRouter()
	router.HandleFunc("/decks", getDecks).Methods("GET")
	router.HandleFunc("/decks/{id}", getDeck).Methods("GET")
	router.HandleFunc("/decks", addDeck).Methods("POST")
	router.HandleFunc("/decks", updateDeck).Methods("PUT")
	router.HandleFunc("/decks/{id}", removeDeck).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getDecks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(decks)
	log.Println("Get decks is called")
}

func getDeck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idInteger, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalln("Could not convert to int")
	}

	for _, deck := range decks {
		if deck.ID == idInteger {
			json.NewEncoder(w).Encode(&deck)
		}
	}
	log.Println("Get deck is called")
}

func addDeck(w http.ResponseWriter, r *http.Request) {
	log.Println("Add deck is called")
}

func updateDeck(w http.ResponseWriter, r *http.Request) {
	log.Println("Update deck is called")
}

func removeDeck(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete deck is called")
}
