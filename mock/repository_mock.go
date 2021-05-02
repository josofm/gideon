package mock

import (
	"github.com/josofm/gideon/model"
)

type RepositoryMock struct {
	User     model.User
	Err      error
	Email    string
	DeckName string
}

func (r *RepositoryMock) Login(email, pass string) (model.User, error) {
	return r.User, r.Err
}

func (r *RepositoryMock) CreateUser(user model.User) (string, error) {
	return r.Email, r.Err
}

func (r *RepositoryMock) GetUser(id uint) (model.User, error) {
	return r.User, r.Err
}

func (r *RepositoryMock) CreateDeck(deck model.Deck) (string, error) {
	return r.DeckName, r.Err
}

func (r *RepositoryMock) UpdateUser(user model.User) error {
	return r.Err
}

func (r *RepositoryMock) DeleteUser(user model.User) error {
	return r.Err
}
