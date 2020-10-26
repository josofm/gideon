package controller

import (
	"errors"
	"log"

	"github.com/josofm/gideon/model"
	repository "github.com/josofm/gideon/repository"
)

type Controller struct {
	repository *repository.Repository
}

func NewController(r *repository.Repository) *Controller {
	return &Controller{
		repository: r,
	}
}

var decks []model.Deck

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Controller) Login(name, pass string) (model.User, error) {
	var user model.User
	if name == "" || pass == "" {
		return user, errors.New("some error temporary")
	}
	user, err := c.repository.Login(name, pass)
	if err != nil {
		return user, err
	}
	return user, nil

}

// func (c Controller) GetDecks(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var deck model.Deck
// 		var error model.Error
// 		decks = []model.Deck{}
// 		deckRepo := repository.DeckRepository{}

// 		decks, err := deckRepo.GetDecks(db, deck, decks)

// 		if err != nil {
// 			error.Message = "Server Error"
// 			utils.SendError(w, http.StatusInternalServerError, error)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		utils.SendSuccess(w, decks)

// 		log.Println("Get decks is called")
// 	}
// }

// func (c Controller) GetDeck(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var deck model.Deck
// 		var error model.Error

// 		decks = []model.Deck{}

// 		deckRepo := repository.DeckRepository{}

// 		params := mux.Vars(r)

// 		id, _ := strconv.Atoi(params["id"])

// 		deck, err := deckRepo.GetDeck(db, deck, id)

// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				error.Message = "Not Found"
// 				utils.SendError(w, http.StatusNotFound, error)
// 				return
// 			} else {
// 				error.Message = "Server error"
// 				utils.SendError(w, http.StatusInternalServerError, error)
// 				return
// 			}
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		utils.SendSuccess(w, deck)

// 		log.Println("Get deck is called")
// 	}
// }

// func (c Controller) AddDeck(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var deck model.Deck
// 		var error model.Error
// 		var deckID int

// 		json.NewDecoder(r.Body).Decode(&deck)

// 		if deck.Commander == "" || deck.Owner == "" {
// 			error.Message = "Enter missing fields."
// 			utils.SendError(w, http.StatusBadRequest, error)
// 			return
// 		}

// 		deckRepo := repository.DeckRepository{}
// 		deckID, err := deckRepo.AddDeck(db, deck)

// 		if err != nil {
// 			error.Message = "Server error"
// 			utils.SendError(w, http.StatusInternalServerError, error)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		utils.SendSuccess(w, deckID)
// 	}

// }

// func (c Controller) UpdateDeck(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var deck model.Deck
// 		var error model.Error

// 		json.NewDecoder(r.Body).Decode(&deck)
// 		if deck.ID == 0 || deck.Owner == "" || deck.Commander == "" {
// 			error.Message = "All fields are required"
// 			utils.SendError(w, http.StatusBadRequest, error)
// 			return
// 		}

// 		deckRepo := repository.DeckRepository{}
// 		rowsUpdated, err := deckRepo.UpdateDeck(db, deck)

// 		if err != nil {
// 			error.Message = "Server error"
// 			utils.SendError(w, http.StatusInternalServerError, error)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "text/plain")
// 		utils.SendSuccess(w, rowsUpdated)

// 	}

// }

// func (c Controller) RemoveDeck(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var error model.Error
// 		params := mux.Vars(r)
// 		deckRepo := repository.DeckRepository{}
// 		id, _ := strconv.Atoi(params["id"])

// 		rowsDeleted, err := deckRepo.RemoveDeck(db, id)

// 		if err != nil {
// 			error.Message = "Server error"
// 			utils.SendError(w, http.StatusInternalServerError, error)
// 			return
// 		}

// 		if rowsDeleted == 0 {
// 			error.Message = "Not Found"
// 			utils.SendError(w, http.StatusInternalServerError, error)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "text/plain")
// 		utils.SendSuccess(w, rowsDeleted)

// 	}
// }
