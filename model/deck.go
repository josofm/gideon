package model

type Deck struct {
	ID        int    `json:id`
	Commander string `json:commander` //will be type card
	Owner     string `json:owner`     //will be type user
}
