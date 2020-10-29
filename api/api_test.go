// +build unit

package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josofm/gideon/api"
	"github.com/josofm/gideon/test"

	"github.com/stretchr/testify/assert"
)

type fixture struct {
	api *api.Api
}

func setup(token map[string]interface{}, err error) fixture {
	c := &test.ControllerMock{}
	c.Token = token
	c.Err = err
	api := api.NewApi(c)

	return fixture{
		api: api,
	}

}

func TestUpAPI(t *testing.T) {
	f := setup(nil, nil)

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

	f := setup(expectedJwt, nil)

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
	f := setup(nil, nil)
	body := []byte(`{"wrongField": "treta"}`)

	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f.api.Login)
	handler.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code Should be equal!")
}
