// +build integration

package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/josofm/gideon/api"
	"github.com/josofm/gideon/clock"
	"github.com/josofm/gideon/controller"
	"github.com/josofm/gideon/repository"
	"github.com/stretchr/testify/assert"
)

const (
	baseUrl               = "http://localhost:80"
	RequestTimeoutSeconds = 180
)

type fixtureIntegration struct {
	a      *api.Api
	client *http.Client
}

//TO DO teardowntest func to clean database and shutdown server
func setupIntegration() fixtureIntegration {

	os.Setenv("ELEPHANTSQL_URL", "postgres://postgres:teste@db:5432/gideondev?sslmode=disable")
	os.Setenv("SECRET_JWT", "secret_token")

	r, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}
	timeout := time.Duration(time.Duration(RequestTimeoutSeconds) * time.Second)
	cli := &http.Client{Timeout: timeout}

	clock := &clock.TimeClock{}
	c := controller.NewController(r, clock)
	a := api.NewApi(c)
	f := fixtureIntegration{
		a:      a,
		client: cli,
	}
	go f.a.StartServer()
	WaitServerUp()
	return f
}

func WaitServerUp() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for range ticker.C {
		res, err := http.Get(baseUrl + "/up")
		if err != nil {
			continue
		}
		if res.StatusCode == http.StatusOK {
			return
		}
	}
}

func TestShouldLoginCorrectly(t *testing.T) {
	_ = setupIntegration()
	body := []byte(`{
	    "email": "gideon@mtg.com",
	    "password": "123change"
	}`)
	resp, err := http.Post(baseUrl+"/login", "aplication/json", bytes.NewBuffer(body))

	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldGetNotFoundWhenLoginNotFound(t *testing.T) {
	_ = setupIntegration()
	body := []byte(`{
	    "email": "yugi@bandai.com",
	    "password": "123change"
	}`)
	resp, err := http.Post(baseUrl+"/login", "aplication/json", bytes.NewBuffer(body))
	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldGetNotFoundPasswordNotMatches(t *testing.T) {
	_ = setupIntegration()
	body := []byte(`{
	    "email": "gideon@mtg.com",
	    "password": "123notchange"
	}`)
	resp, err := http.Post(baseUrl+"/login", "aplication/json", bytes.NewBuffer(body))

	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldRegisterCorrectly(t *testing.T) {
	_ = setupIntegration()
	body := []byte(`{
	    "email": "liliana@mtg.com",
	    "password": "123change",
	    "sex": "f",
	    "age": "32",
	    "name": "Liliana Vess"
	}`)
	resp, err := http.Post(baseUrl+"/register", "aplication/json", bytes.NewBuffer(body))

	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

//In this case, is a sequential test
func TestShouldGetErrorWhenEmailAlreadyRegister(t *testing.T) {
	_ = setupIntegration()
	body := []byte(`{
	    "email": "gideon@mtg.com",
	    "password": "123change",
	    "sex": "m",
	    "age": "32",
	    "name": "Gideon Jura"
	}`)
	resp, err := http.Post(baseUrl+"/register", "aplication/json", bytes.NewBuffer(body))

	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusConflict, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldGetErrorWhenReceiveAWrongBody(t *testing.T) {
	_ = setupIntegration()
	body := []byte(`{
	    "email": "liliana@mtg.com",
	    "age": "32",
	    "name": "Liliana Vess"
	}`)
	resp, err := http.Post(baseUrl+"/register", "aplication/json", bytes.NewBuffer(body))

	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldGetUserCorrectly(t *testing.T) {
	f := setupIntegration()

	body := []byte(`{
	    "email": "gideon@mtg.com",
	    "password": "123change"
	}`)
	token := loginTest(body)

	r, err := http.NewRequest("GET", baseUrl+"/auth/user/1", nil)
	r.Header.Add("access-token", token["token"].(string))
	resp, err := f.client.Do(r)
	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")

}

func TestShouldGetForbiddenWhenUserIdDidNotMatches(t *testing.T) {
	f := setupIntegration()

	body := []byte(`{
	    "email": "gideon@mtg.com",
	    "password": "123change"
	}`)
	token := loginTest(body)

	r, err := http.NewRequest("GET", baseUrl+"/auth/user/7", nil)
	r.Header.Add("access-token", token["token"].(string))
	resp, err := f.client.Do(r)
	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusForbidden, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldUpdateUserCorrectly(t *testing.T) {
	f := setupIntegration()

	body := []byte(`{
	    "email": "tibalt@mtg.com",
	    "password": "123change"
	}`)

	token := loginTest(body)
	r, err := http.NewRequest("DELETE", baseUrl+"/auth/user/2", nil)
	r.Header.Add("access-token", token["token"].(string))
	resp, err := f.client.Do(r)

	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")

}

func TestGetCardNameCorrectly(t *testing.T) {
	f := setupIntegration()
	r, err := http.NewRequest("GET", baseUrl+"/card/hogaak", nil)
	resp, err := f.client.Do(r)
	assert.NotNil(t, resp.Body, "Should be not nil!")
	assert.Nil(t, err, "Should be nil!")
}

func loginTest(body []byte) map[string]interface{} {
	loginRequest, _ := http.Post(baseUrl+"/login", "aplication/json", bytes.NewBuffer(body))

	var token map[string]interface{}
	decoder := json.NewDecoder(loginRequest.Body)
	_ = decoder.Decode(&token)

	return token
}
