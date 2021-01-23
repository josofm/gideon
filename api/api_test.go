// +build unit

package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/josofm/gideon/api"
	"github.com/josofm/gideon/mock"
	"github.com/josofm/gideon/model"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	api *api.Api
	r   *mux.Router
}

func setup(c *mock.ControllerMock) fixture {
	api := api.NewApi(c)

	router := mux.NewRouter().StrictSlash(true)
	router.Use()
	router.HandleFunc("/up", api.Up).Methods("GET")
	router.HandleFunc("/login", api.Login).Methods("POST")
	router.HandleFunc("/register", api.Register).Methods("POST")
	router.HandleFunc("/card/{name}", api.GetCardByName).Methods("GET")
	s := router.PathPrefix("/auth").Subrouter()
	s.Use(api.JwtVerify)
	s.HandleFunc("/user/{id}", api.GetUser).Methods("GET")
	s.HandleFunc("/deck", api.AddDeck).Methods("POST")

	return fixture{
		api: api,
		r:   router,
	}

}

func TestUpAPI(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = nil
	c.Email = ""
	f := setup(c)

	r, err := http.NewRequest("GET", "/up", nil)

	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")

}

func TestShouldLoginCorrectly(t *testing.T) {
	expectedJwt := map[string]interface{}{
		"token": "tokenzera",
	}
	c := &mock.ControllerMock{}
	c.Token = expectedJwt
	c.Err = nil
	c.Email = ""
	f := setup(c)

	body := []byte(`{"email": "gideon@mtg.com", "password": "ravnica"}`)

	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	var actual map[string]interface{}
	_ = json.Unmarshal(rr.Body.Bytes(), &actual)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
	assert.Equal(t, expectedJwt, actual, "Token should be equal!")

}

func TestShouldGetErrorWhenBodyIsWrong(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = nil
	c.Email = ""
	f := setup(c)

	body := []byte(`{"wrongField": "treta"}`)

	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}

func TestShouldGetErrorWhenInvalidCredentials(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = errors.New("login invalid")
	c.Email = ""
	f := setup(c)

	body := []byte(`{"email": "gideon@mtg.com", "password": "ravnica"}`)

	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusNotFound, rr.Code, "Status code Should be equal!")
}

func TestShouldRegisterNewUserCorrectly(t *testing.T) {
	expectedEmail := "capeta@mtg.com"
	ExpectedMessage := fmt.Sprintf("Welcome %v", expectedEmail)

	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = nil
	c.Email = expectedEmail
	f := setup(c)
	body := []byte(`
		{
			"name": "capeta da charneca",
			"sex": "m",
			"age": "10000",
			"password": "fatherbolas",
			"email": "capeta@mtg.com"
		}
	`)
	r, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	var email string
	_ = json.Unmarshal(rr.Body.Bytes(), &email)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
	assert.Equal(t, ExpectedMessage, email, "Token should be equal!")

}

func TestShouldGetErrorToRegisterWhenUserInformWrongBody(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = nil
	c.Email = ""
	f := setup(c)
	body := []byte(`{"wrongField": "treta"}`)

	r, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}

func TestShouldGetErrorToRegisterWhenUserInformFewValues(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = errors.New("missing parameters")
	c.Email = ""
	f := setup(c)

	body := []byte(`
		{
			"name": "capeta da charneca",
			"sex": "m"
		}`)

	r, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}

func TestShouldGetErrorWhenNotInformToken(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = nil
	c.Email = ""
	f := setup(c)

	body := []byte(`{"nottoken": ""}`)

	r, err := http.NewRequest("GET", "/auth/user/32", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)
	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusForbidden, rr.Code, "Status code Should be equal!")
}

func TestShouldGetErrorWhenCantParseTheToken(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.Err = errors.New("generic error")
	c.Email = ""
	f := setup(c)

	body := []byte(``)
	r, err := http.NewRequest("GET", "/auth/user/32", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)
	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusForbidden, rr.Code, "Status code Should be equal!")
}

func TestShouldGetStatusForbiddenWhenTryGetUserWithDidNotmatchIds(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.ErrGetUser = nil
	c.Email = ""
	c.TokenModel = model.Token{
		UserID: 1,
		Name:   "capeta da charneca",
		Email:  "swamp@house.com",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	f := setup(c)

	body := []byte(``)
	r, err := http.NewRequest("GET", "/auth/user/s", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusForbidden, rr.Code, "Status code Should be equal!")
}

func TestShouldGetStatusForbiddenWhenTryGetUserWithNotValidParameter(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.ErrGetUser = errors.New("Error getting user")
	c.Email = ""
	c.TokenModel = model.Token{
		UserID: 1,
		Name:   "capeta da charneca",
		Email:  "swamp@house.com",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	f := setup(c)

	body := []byte(``)
	r, err := http.NewRequest("GET", "/auth/user/1", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusNotFound, rr.Code, "Status code Should be equal!")
}

func TestShouldGetUserCorrectly(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.ErrGetUser = nil
	c.Email = ""
	c.TokenModel = model.Token{
		UserID: 1,
		Name:   "capeta da charneca",
		Email:  "swamp@house.com",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	f := setup(c)

	body := []byte(``)
	r, err := http.NewRequest("GET", "/auth/user/1", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
}

func TestShouldGetBadRequestWhenMalformedToken(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.ErrGetUser = nil
	c.Email = ""

	f := setup(c)
	body := []byte(``)
	r, err := http.NewRequest("GET", "/auth/user/1", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code Should be equal!")
}

func TestShouldGetCardByIdCorrectly(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Cards = []*mtg.Card{
		&mtg.Card{
			Name: "Hogaak, Arisen Necropolis",
		},
	}
	f := setup(c)
	body := []byte(``)
	r, err := http.NewRequest("GET", "/card/hoogak", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	var cards []*mtg.Card
	_ = json.Unmarshal(rr.Body.Bytes(), &cards)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, len(c.Cards), len(cards), "Should be Equal!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
}

func TestShouldInsertDeckCorrectly(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.ErrGetUser = nil
	c.Email = ""
	c.ErrGetDeck = nil
	c.DeckName = "best deck"
	c.TokenModel = model.Token{
		UserID: 1,
		Name:   "capeta da charneca",
		Email:  "swamp@house.com",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	f := setup(c)
	body := []byte(`{
		"name": "best deck",
		"commander": {
			"card": {
				"multiverseID": 389712
			}
		},
		"cards": [
			{
				"card": {
					"multiverseID": 194969
				}
			}
		]
	}`)
	r, err := http.NewRequest("POST", "/auth/deck", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
}

func TestShouldGetErrorInsertingADeckWhenBodyDidNotHaveInformation(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.ErrGetUser = nil
	c.Email = ""
	c.TokenModel = model.Token{
		UserID: 1,
		Name:   "capeta da charneca",
		Email:  "swamp@house.com",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	f := setup(c)
	body := []byte(``)
	r, err := http.NewRequest("POST", "/auth/deck", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}

func TestShouldInsertDeckCorrectlyButGetErrorBecauseIsNotValidCommander(t *testing.T) {
	c := &mock.ControllerMock{}
	c.Token = nil
	c.ErrGetUser = nil
	c.Email = ""
	c.ErrGetDeck = errors.New("A commander must be a legendary card")
	c.DeckName = ""
	c.TokenModel = model.Token{
		UserID: 1,
		Name:   "capeta da charneca",
		Email:  "swamp@house.com",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	f := setup(c)
	body := []byte(`{
		"name": "best deck",
		"commander": {
			"card": {
				"multiverseID": 409574
			}
		},
		"cards": [
			{
				"card": {
					"multiverseID": 194969
				}
			}
		]
	}`)
	r, err := http.NewRequest("POST", "/auth/deck", bytes.NewBuffer(body))
	r.Header.Set("access-token", "humansoldier1/1")
	rr := httptest.NewRecorder()
	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}
