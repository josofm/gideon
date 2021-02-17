package user

import (
	"errors"
	"log"

	"github.com/josofm/gideon/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (u *UserRepository) Login(email, pass string, dbPool *gorm.DB) (model.User, error) {
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

func (u *UserRepository) getUserByEmail(email string, dbPool *gorm.DB) (model.User, error) {
	user := model.User{}
	if result := dbPool.Where("email = ?", email).First(&user); result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (u *UserRepository) Create(user model.User, dbPool *gorm.DB) (string, error) {
	_, err := u.getUserByEmail(user.Email, dbPool)
	if err == nil {
		log.Print("[Create] This user already exists")
		return "", errors.New("User already Registred")
	} else if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("[Create] Some sql problem")
		return "", err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("[Create User] Fail encripting")
		return "", err
	}
	user.Password = string(pass)
	if result := dbPool.Create(&user); result.Error != nil {
		log.Print("[Create User] Fail database insertion")
		return "", result.Error

	}
	return user.Email, nil
}

func (u *UserRepository) Get(id uint, dbPool *gorm.DB) (model.User, error) {
	user := model.User{}
	if result := dbPool.Where("ID = ?", id).First(&user); result.Error != nil {
		log.Print("[Get] User not found in database")
		return model.User{}, result.Error
	}
	return user, nil
}

func (u *UserRepository) Update(user model.User, dbPool *gorm.DB) error {
	if _, err := u.Get(user.ID, dbPool); err != nil {
		return err
	}
	if result := dbPool.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) Delete(user model.User, dbPool *gorm.DB) error {
	if result := dbPool.Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
