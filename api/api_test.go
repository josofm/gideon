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

	"github.com/josofm/gideon/api"
	"github.com/josofm/gideon/mock"

	"github.com/stretchr/testify/assert"
)

type fixture struct {
	api *api.Api
}

func setup(token map[string]interface{}, err error, email string) fixture {
	c := &mock.ControllerMock{}
	c.Token = token
	c.Err = err
	c.Email = email
	api := api.NewApi(c)

	return fixture{
		api: api,
	}

}

func TestUpAPI(t *testing.T) {
	f := setup(nil, nil, "")

	handler := http.HandlerFunc(f.api.Up)
	r, err := http.NewRequest("GET", "/up", nil)

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")

}

func TestShouldLoginCorrectly(t *testing.T) {
	expectedJwt := map[string]interface{}{
		"token": "tokenzera",
	}

	f := setup(expectedJwt, nil, "")

	body := []byte(`{"email": "gideon@mtg.com", "password": "ravnica"}`)

	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f.api.Login)

	handler.ServeHTTP(rr, r)

	var actual map[string]interface{}
	_ = json.Unmarshal(rr.Body.Bytes(), &actual)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
	assert.Equal(t, expectedJwt, actual, "Token should be equal!")

}

func TestShouldGetErrorWhenBodyIsWrong(t *testing.T) {
	f := setup(nil, nil, "")
	body := []byte(`{"wrongField": "treta"}`)

	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f.api.Login)
	handler.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}

func TestShouldGetErrorWhenInvalidCredentials(t *testing.T) {
	f := setup(nil, errors.New("login invalid"), "")
	body := []byte(`{"email": "gideon@mtg.com", "password": "ravnica"}`)

	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f.api.Login)
	handler.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusNotFound, rr.Code, "Status code Should be equal!")
}

func TestShouldRegisterNewUserCorrectly(t *testing.T) {
	expectedEmail := "capeta@mtg.com"
	ExpectedMessage := fmt.Sprintf("Welcome %v", expectedEmail)
	f := setup(nil, nil, expectedEmail)
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
	handler := http.HandlerFunc(f.api.Register)
	handler.ServeHTTP(rr, r)

	var email string
	_ = json.Unmarshal(rr.Body.Bytes(), &email)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
	assert.Equal(t, ExpectedMessage, email, "Token should be equal!")

}

func TestShouldGetErrorToRegisterWhenUserInformWrongBody(t *testing.T) {
	f := setup(nil, nil, "")
	body := []byte(`{"wrongField": "treta"}`)

	r, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f.api.Register)
	handler.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}

func TestShouldGetErrorToRegisterWhenUserInformFewValues(t *testing.T) {
	f := setup(nil, errors.New("missing parameters"), "")
	body := []byte(`
		{
			"name": "capeta da charneca",
			"sex": "m"
		}`)

	r, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f.api.Register)
	handler.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be nil!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}
