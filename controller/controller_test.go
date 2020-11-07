// +build unit

package controller_test

import (
	"errors"
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

func setup(user model.User, err error) fixture {
	os.Setenv("SECRET_JWT", "generic_jwt_token")
	r := &mock.RespositoryMock{}
	r.User = user
	r.Err = err

	clock := &mock.ClockMock{}
	clock.NowMock = time.Date(2009, 04, 30, 20, 34, 58, 651387237, time.UTC)

	c := controller.NewController(r, clock)

	return fixture{
		c: c,
	}

}

func TestShouldGetTokenLoginCorrectly(t *testing.T) {
	u := model.User{
		ID:       1.0,
		Name:     "jace belerem",
		Sex:      "m",
		Age:      "12",
		Email:    "jace@mtg.com",
		Password: "$3dsfTrcsa",
	}
	f := setup(u, nil)
	token, err := f.c.Login(u.Email, u.Password)
	expectedToken := map[string]interface{}{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiTmFtZSI6ImphY2UgYmVsZXJlbSIsIkVtYWlsIjoiamFjZUBtdGcuY29tIiwiZXhwIjoxMjQ3MTIzNjk4fQ.R2d8MZAeOVaX1Qs23UPoM5zHDd8YqNqAhc6y2G0Fvu8",
	}

	assert.Nil(t, err, "should be nil!")
	assert.Equal(t, expectedToken, token, "Should be equal!")
}

func TestShouldGetErrorWhenCantLogin(t *testing.T) {
	f := setup(model.User{}, errors.New("Invalid Credentials"))
	token, err := f.c.Login("ComboPlayer@mtg.com", "123Change")

	assert.NotNil(t, err, "should be not nil!")
	assert.Nil(t, token, "Should be nil!")
}
