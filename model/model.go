package model

import (
	"github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/dgrijalva/jwt-go"
)

type Deck struct {
	ID        float64 `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Commander Card    `json:commander,omitempty`
	Owner     User    `json:owner,omitempty`
	Cards     []Card  `json:"cards,omitempty"`
}

type Card struct {
	Card  mtg.Card `json:"card,omitempty"`
	Price Price    `json:"price, omitempty"`
}

type Price struct {
	Minimun float64 `json:"minimun,omitempty"`
	Average float64 `json:"average,omitempty"`
	Maximum float64 `json:"maximun,omitempty"`
}

type User struct {
	ID       float64 `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Sex      string  `json:"sex,omitempty"`
	Age      string  `json:"age,omitempty"`
	Token    string  `json:"token,omitempty"`
	Email    string  `json:"email,omitempty"`
	Password string  `json:"password,omitempty"`
}

type Token struct {
	UserID float64
	Name   string
	Email  string
	*jwt.StandardClaims
}
