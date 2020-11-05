package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/josofm/gideon/model"
	"github.com/josofm/gideon/repository/deck"
	"github.com/josofm/gideon/repository/user"

	"github.com/lib/pq"
)

type Repository struct {
	url    string
	user   *user.UserRepository
	deck   *deck.DeckRepository
	dbPool *sql.DB
}

func NewRepository() (*Repository, error) {
	r := &Repository{}
	ur := user.NewUserRepository()
	dr := deck.NewDeckRepository()
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Print("Error parsing sql url")
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

	r.dbPool, err = sql.Open("postgres", r.url)
	if err != nil {
		return err
	}

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

// func (r *Repository) Login(name, pass string) (model.User, error) {
// 	user := model.User{}
// 	rows := db.Query("select * from user where name=$1 and password=$2", name, pass)
// 	err := rows.Scan(&user.ID, &user.Name, &user.Sex, &user.Age, &user.Token)
// 	if err != nil {
// 		return model.User{}, err
// 	}

// 	return model.User, nil
// }
