package controller

import (
	"github.com/josofm/gideon/model"
)

func userHasAllFields(user model.User) bool {
	if user.Name == "" {
		return false
	} else if user.Sex == "" {
		return false
	} else if user.Age == "" {
		return false
	} else if user.Email == "" {
		return false
	} else if user.Password == "" {
		return false
	}
	return true
}
