package user

import (
	"database/sql"

	"github.com/josofm/gideon/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Login(name, pass string, dbPool *sql.DB) (model.User, error) {
	user := model.User{}
	rows := dbPool.QueryRow("select id from user where name=$1 and password=$2", name, pass)
	err := rows.Scan(&user.ID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
