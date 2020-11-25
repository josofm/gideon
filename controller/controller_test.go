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
	r := &mock.RepositoryMock{}
	r.User = user
	r.Err = err
	r.Email = user.Email

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
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiJqYWNlIGJlbGVyZW0iLCJFbWFpbCI6ImphY2VAbXRnLmNvbSIsImV4cCI6MTI0NzEyMzY5OH0.gcsOKKUBXeqrvPEoizoUL-Gk2BH38QT7EU4hM81XHJs",
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

func TestShouldRegisterNewUserCorrectly(t *testing.T) {
	u := model.User{
		Name:     "jace belerem",
		Sex:      "m",
		Age:      "12",
		Email:    "jace@mtg.com",
		Password: "$3dsfTrcsa",
	}
	f := setup(u, nil)
	email, err := f.c.CreateUser(u)
	assert.Nil(t, err, "should be nil!")
	assert.Equal(t, u.Email, email, "Should be equal!")
}

func TestShouldGetErrorWhenMissingFields(t *testing.T) {
	u := model.User{
		Name: "jace belerem",
		Sex:  "m",
		Age:  "12",
	}
	f := setup(u, nil)
	email, err := f.c.CreateUser(u)

	assert.NotNil(t, err, "should be nil!")
	assert.Equal(t, "", email, "Should be equal!")
}

func TestShouldGetErrorWhenRegisterError(t *testing.T) {
	u := model.User{
		Name:     "jace belerem",
		Sex:      "m",
		Age:      "12",
		Email:    "jace@mtg.com",
		Password: "$3dsfTrcsa",
	}
	f := setup(u, errors.New("missing fields"))
	email, err := f.c.CreateUser(u)

	assert.NotNil(t, err, "should be nil!")
	assert.Equal(t, "", email, "Should be equal!")
}
