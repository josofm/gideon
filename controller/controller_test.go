// +build unit

package controller_test

import (
	"os"
	"testing"
	"time"

	"github.com/josofm/gideon/controller"
	"github.com/josofm/gideon/mock"
	"github.com/josofm/gideon/model"

	"github.com/stretchr/testify/assert"
)

type fixture struct {
	c *controller.Controller
}

func setup(user model.User, err error, deckName string) fixture {
	os.Setenv("SECRET_JWT", "generic_jwt_token")
	r := &mock.RepositoryMock{}
	r.User = user
	r.Err = err
	r.Email = user.Email
	r.DeckName = deckName

	clock := &mock.ClockMock{}
	clock.NowMock = time.Date(2020, 04, 30, 20, 34, 58, 651387237, time.UTC)

	c := controller.NewController(r, clock)

	return fixture{
		c: c,
	}

}

// func TestShouldGetTokenLoginCorrectly(t *testing.T) {
// 	u := model.User{
// 		ID:       1,
// 		Name:     "jace belerem",
// 		Sex:      "m",
// 		Age:      "12",
// 		Email:    "jace@mtg.com",
// 		Password: "$3dsfTrcsa",
// 	}
// 	f := setup(u, nil, "")
// 	token, err := f.c.Login(u.Email, u.Password)
// 	expectedToken := map[string]interface{}{
// 		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiJqYWNlIGJlbGVyZW0iLCJFbWFpbCI6ImphY2VAbXRnLmNvbSIsImV4cCI6MTU4ODI4MjQ5OH0.GZM0n5fECxhXZl-_r37q8tapS8i2xQIp_v9t6UoTcz0",
// 	}

// 	assert.Nil(t, err, "should be nil!")
// 	assert.Equal(t, expectedToken, token, "Should be equal!")
// }

// func TestShouldGetErrorWhenCantLogin(t *testing.T) {
// 	f := setup(model.User{}, errors.New("Invalid Credentials"), "")
// 	token, err := f.c.Login("ComboPlayer@mtg.com", "123Change")

// 	assert.NotNil(t, err, "should be not nil!")
// 	assert.Nil(t, token, "Should be nil!")
// }

// func TestShouldRegisterNewUserCorrectly(t *testing.T) {
// 	u := model.User{
// 		Name:     "jace belerem",
// 		Sex:      "m",
// 		Age:      "12",
// 		Email:    "jace@mtg.com",
// 		Password: "$3dsfTrcsa",
// 	}
// 	f := setup(u, nil, "")
// 	email, err := f.c.CreateUser(u)
// 	assert.Nil(t, err, "should be nil!")
// 	assert.Equal(t, u.Email, email, "Should be equal!")
// }

// func TestShouldGetErrorWhenMissingFields(t *testing.T) {
// 	u := model.User{
// 		Name: "jace belerem",
// 		Sex:  "m",
// 		Age:  "12",
// 	}
// 	f := setup(u, nil, "")
// 	email, err := f.c.CreateUser(u)

// 	assert.NotNil(t, err, "should be nil!")
// 	assert.Equal(t, "", email, "Should be equal!")
// }

// func TestShouldGetErrorWhenRegisterError(t *testing.T) {
// 	u := model.User{
// 		Name:     "jace belerem",
// 		Sex:      "m",
// 		Age:      "12",
// 		Email:    "jace@mtg.com",
// 		Password: "$3dsfTrcsa",
// 	}
// 	f := setup(u, errors.New("missing fields"), "")
// 	email, err := f.c.CreateUser(u)

// 	assert.NotNil(t, err, "should be nil!")
// 	assert.Equal(t, "", email, "Should be equal!")
// }

// func TestShouldGetErrorWhenTryParseToken(t *testing.T) {
// 	f := setup(model.User{}, nil, "")
// 	tk, err := f.c.GetToken("tokenzeraWrong")
// 	assert.NotNil(t, err, "should be nil!")
// 	assert.Equal(t, model.Token{}, tk, "Should be equal!")
// }

// func TestShouldGetErrorWhenTryGetUser(t *testing.T) {
// 	f := setup(model.User{}, errors.New("User not found"), "")
// 	user, err := f.c.GetUser(2)
// 	assert.NotNil(t, err, "should be nil!")
// 	assert.Equal(t, model.User{}, user, "Should be equal!")
// }

// func TestShouldGetUserCorrectly(t *testing.T) {
// 	u := model.User{
// 		Name:     "jace belerem",
// 		Sex:      "m",
// 		Age:      "12",
// 		Email:    "jace@mtg.com",
// 		Password: "$3dsfTrcsa",
// 	}
// 	f := setup(u, nil, "")
// 	user, err := f.c.GetUser(2)
// 	assert.Nil(t, err, "should be nil!")
// 	assert.Equal(t, u, user, "Should be equal!")
// }

func TestShouldGetCardByNameCorrectly(t *testing.T) {
	f := setup(model.User{}, nil, "")
	cards, err := f.c.GetCardByName("hogaak")

	expectedSize := 1

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, expectedSize, len(cards), "Should be Equal!")
}

func TestShouldCreateDeckCorrectly(t *testing.T) {
	expectedName := "best deck"
	f := setup(model.User{}, nil, expectedName)
	deck := mock.GetBasicDeck()
	deckName, err := f.c.CreateDeck(deck, 1)
	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, expectedName, deckName, "Should be Equal!")
}

// func TestShouldGetErrorCreatingADeckWhenCardCantBeCommander(t *testing.T) {
// 	expectedName := ""
// 	f := setup(model.User{}, nil, expectedName)
// 	deck := mock.GetBasicDeck()
// 	deck.Commander.Card.MultiverseId = 409574
// 	deckName, err := f.c.CreateDeck(deck, 1)
// 	assert.NotNil(t, err, "Should be nil!")
// 	assert.Equal(t, expectedName, deckName, "Should be Equal!")
// }

// func TestShouldGetErrorCreatingADeckWhenCardisBanned(t *testing.T) {
// 	expectedName := ""
// 	f := setup(model.User{}, nil, expectedName)
// 	deck := mock.GetBasicDeck()
// 	deck.Commander.Card.MultiverseId = 425897
// 	deckName, err := f.c.CreateDeck(deck, 1)
// 	assert.NotNil(t, err, "Should be nil!")
// 	assert.Equal(t, expectedName, deckName, "Should be Equal!")
// }

// func TestShouldDeleteUserCorrectly(t *testing.T) {
// 	expectedName := ""
// 	f := setup(model.User{}, nil, expectedName)
// 	err := f.c.DeleteUser(1)
// 	assert.Nil(t, err, "Should be nil!")
// }

// func TestShouldGetErrorDeleteingUser(t *testing.T) {
// 	f := setup(model.User{}, errors.New("User not found"), "")
// 	err := f.c.DeleteUser(1)
// 	assert.NotNil(t, err, "Should be not nil!")
// }
