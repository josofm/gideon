package test

import (
	"testing"

	"github.com/josofm/gideon/model"
	"github.com/josofm/gideon/repository"
	"github.com/stretchr/testify/assert"
)

func Clean(r *repository.Repository) {
	r.Clean("test")
}

func CreateUser(r *repository.Repository, user model.User, t *testing.T) {
	_, err := r.CreateUser(user)
	assert.Nil(t, err, "Should be nil!")
}
