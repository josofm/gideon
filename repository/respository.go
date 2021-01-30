package repository

import (
	"database/sql"
	"os"

	"github.com/josofm/gideon/model"
	"github.com/josofm/gideon/repository/deck"
	"github.com/josofm/gideon/repository/user"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Repository struct {
	url    string
	user   *user.UserRepository
	deck   *deck.DeckRepository
	dbPool *gotm.DB
}

func NewRepository() (*Repository, error) {
	r := &Repository{}
	ur := &user.UserRepository{}
	dr := &deck.DeckRepository{}
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		return r, err
	}
	r.url = pgUrl
	r.deck = dr
	r.user = ur
	return r, nil
}

func (r *Repository) connectDB() error {
	var err error
	if r.dbPool != nil {
		return nil
	}
	r.dbPool, err = gorm.Open(sql.Open("postgres", r.url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// r.dbPool, err = sql.Open("postgres", r.url)
	// if err != nil {
	// 	return err
	// }

	err = r.dbPool.Ping()
	if err != nil {
		return err
	}

	return nil

}

func (r *Repository) Login(email, pass string) (model.User, error) {
	err := r.connectDB()
	if err != nil {
		return model.User{}, err
	}
	return r.user.Login(email, pass, r.dbPool)
}

func (r *Repository) CreateUser(user model.User) (string, error) {
	err := r.connectDB()
	if err != nil {
		return "", err
	}
	return r.user.Create(user, r.dbPool)
}

func (r *Repository) GetUser(id float64) (model.User, error) {
	err := r.connectDB()
	if err != nil {
		return model.User{}, err
	}
	return r.user.Get(id, r.dbPool)
}

func (r *Repository) CreateDeck(deck model.Deck) (string, error) {
	err := r.connectDB()
	if err != nil {
		return "", err
	}
	return r.deck.Create(deck, r.dbPool)
}
