// +build integration

package api_test

import (
	"bytes"
	"fmt"
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
	baseUrl = "http://localhost:80"
)

type fixtureIntegration struct {
	a *api.Api
}

func setupIntegration() fixtureIntegration {

	os.Setenv("ELEPHANTSQL_URL", "postgres://postgres:teste@localhost:5432/gideondev")

	r, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}
	clock := &clock.TimeClock{}
	c := controller.NewController(r, clock)
	a := api.NewApi(c)
	f := fixtureIntegration{
		a: a,
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

func TestWrong(t *testing.T) {
	_ = setupIntegration()
	body := []byte(`{
	    "email": "gideon@mtg.com",
	    "password": "123change"
	}`)
	resp, err := http.Post(baseUrl+"/login", "aplication/json", bytes.NewBuffer(body))
	fmt.Println(resp)
	fmt.Println(err)

	assert.True(t, false, "test")
}
