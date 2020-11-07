package user

import (
	"database/sql"
	"log"

	"github.com/josofm/gideon/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (u *UserRepository) Login(email, pass string, dbPool *sql.DB) (model.User, error) {
	user := model.User{}
	rows := dbPool.QueryRow("select * from user where email=$1", email)
	err := rows.Scan(&user)
	if err != nil {
		return model.User{}, err
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		log.Printf("Invalid login credentials. Please try again - %v", errf)
		return model.User{}, errf
	}

	return user, nil
}
