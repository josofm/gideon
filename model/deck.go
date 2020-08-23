package model

import (
	"github.com/MagicTheGathering/mtg-sdk-go"
)

type Deck struct {
	ID        int    `json:id,omitempty`
	Commander string `json:commander,omitempty` //will be type card
	//Commander mtg.Card `json:commander,omitempty`
	//Owner     User   `json:owner,omitempty`
	Owner string     `json:owner,omitempty` //will be type user
	Cards []mtg.Card `json:cards,omitempty`
}

type User struct {
	ID   int    `json:id,omitempty`
	Name string `json:name,omitempty`
}
