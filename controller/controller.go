package controller

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/josofm/gideon/model"
)

type Controller struct {
	repository Repository
	clock      TimeClock
	secret     string
}

type Repository interface {
	Login(email, pass string) (model.User, error)
	CreateUser(user model.User) (string, error)
	GetUser(id float64) (model.User, error)
}

type TimeClock interface {
	Now() time.Time
	Add(t time.Time, d time.Duration) time.Time
}

func NewController(r Repository, c TimeClock) *Controller {
	s := os.Getenv("SECRET_JWT")
	return &Controller{
		repository: r,
		clock:      c,
		secret:     s,
	}
}

func (c *Controller) Login(email, pass string) (map[string]interface{}, error) {
	var user model.User
	if email == "" || pass == "" {
		return nil, errors.New("some error temporary")
	}
	user, err := c.repository.Login(email, pass)
	if err != nil {
		return nil, err
	}

	t, err := c.createToken(user)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (c *Controller) createToken(user model.User) (map[string]interface{}, error) {
	expiresAt := c.clock.Now().Add(time.Hour * 1).Unix()
	tk := &model.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(c.secret))
	if err != nil {
		return nil, err
	}
	t := map[string]interface{}{
		"token": tokenString,
	}
	return t, nil
}

func (c *Controller) CreateUser(user model.User) (string, error) {
	if !userHasAllFields(user) {
		log.Print("[createUser] missing parameters")
		return "", errors.New("Missing parameters")
	}
	email, err := c.repository.CreateUser(user)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (c *Controller) GetToken(header string) (model.Token, error) {
	tk := &model.Token{}

	_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.secret), nil
	})
	if err != nil {
		return model.Token{}, err
	}
	return *tk, err
}

func (c *Controller) GetUser(id float64) (model.User, error) {
	var user model.User
	user, err := c.repository.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil

}

func userHasAllFields(user model.User) bool {
	if user.Name == "" {
		return false
	} else if user.Sex == "" {
		return false
	} else if user.Age == "" {
		return false
	} else if user.Email == "" {
		return false
	} else if user.Password == "" {
		return false
	}
	return true
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
