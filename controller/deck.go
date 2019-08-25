package controller

import (
	"commander-list/model"
	"commander-list/repository/deck"
	"commander-list/utils"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Controller struct{}

var decks []model.Deck

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetDecks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deck model.Deck
		var error model.Error
		decks = []model.Deck{}
		deckRepo := deckRepository.DeckRepository{}

		decks, err := deckRepo.GetDecks(db, deck, decks)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, decks)

		log.Println("Get decks is called")
	}
}

func (c Controller) GetDeck(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deck model.Deck
		var error model.Error

		decks = []model.Deck{}

		deckRepo := deckRepository.DeckRepository{}

		params := mux.Vars(r)

		id, _ := strconv.Atoi(params["id"])

		deck, err := deckRepo.GetDeck(db, deck, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, deck)

		log.Println("Get deck is called")
	}
}
