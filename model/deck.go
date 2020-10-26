package model

import (
	"github.com/MagicTheGathering/mtg-sdk-go"
)

type Deck struct {
	ID        int    `json:"id,omitempty"`
	Commander string `json:"commander,omitempty"` //will be type car"d
	//Commander mtg.Card `json:commander,omitempty`
	//Owner     User   `json:owner,omitempty`
	Owner string     `json:"owner,omitempty"` //will be type user
	Cards []mtg.Card `json:"cards,omitempty"`
}

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Age      string `json:"age,omitempty"`
	Token    string `json:"token,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
