package user

import (
	"database/sql"
	"errors"
	"log"

	"github.com/josofm/gideon/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

//TO DO use gorm
func (u *UserRepository) Login(email, pass string, dbPool *sql.DB) (model.User, error) {
	user, err := u.getUserByEmail(email, dbPool)
	if err != nil || (model.User{}) == user {
		log.Print("[Login] email not found")
		return model.User{}, err
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		log.Printf("Invalid login credentials. Please try again - %v", errf)
		return model.User{}, errf
	}

	return user, nil
}

func (u *UserRepository) getUserByEmail(email string, dbPool *sql.DB) (model.User, error) {
	user := model.User{}
	rows := dbPool.QueryRow(`select * from "user" as u where u.email=$1`, email)
	err := rows.Scan(&user.ID, &user.Name, &user.Sex, &user.Age, &user.Password, &user.Email)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepository) Create(user model.User, dbPool *sql.DB) (string, error) {
	_, err := u.getUserByEmail(user.Email, dbPool)
	if err == nil {
		log.Print("[Create] This user already exists")
		return "", errors.New("User already Registred")
	} else if err != nil && err != sql.ErrNoRows {
		log.Print("[Create] Some sql problem")
		return "", err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("[Create User] Fail encripting")
		return "", err
	}
	user.Password = string(pass)
	insertStatment := `INSERT INTO "user" (name,sex,age,email,password) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err = dbPool.QueryRow(insertStatment, user.Name, user.Sex, user.Age, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Print("[Create User] Fail database insertion")
		return "", err
	}

	return user.Email, nil
}

func (u *UserRepository) Get(id float64, dbPool *sql.DB) (model.User, error) {
	user := model.User{}
	rows := dbPool.QueryRow(`select * from "user" as u where u.id=$1`, id)
	err := rows.Scan(&user.ID, &user.Name, &user.Sex, &user.Age, &user.Password, &user.Email)
	if err != nil {
		log.Print("[Get] User not found in database")
		return model.User{}, err
	}
	return user, nil
}
