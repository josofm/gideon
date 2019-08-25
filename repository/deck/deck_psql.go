package deckRepository

import (
	"commander-list/model"
	"database/sql"
)

type DeckRepository struct{}

func (d DeckRepository) GetDecks(db *sql.DB, deck model.Deck, decks []model.Deck) ([]model.Deck, error) {

	rows, err := db.Query("select * from decks")
	if err != nil {
		return []model.Deck{}, err
	}

	for rows.Next() {
		err = rows.Scan(&deck.ID, &deck.Commander, &deck.Owner)

		decks = append(decks, deck)
	}

	if err != nil {
		return []model.Deck{}, err
	}

	return decks, nil

}

func (d DeckRepository) GetDeck(db *sql.DB, deck model.Deck, id int) (model.Deck, error) {
	rows := db.QueryRow("select * from decks where id=$1", id)
	err := rows.Scan(&deck.ID, &deck.Commander, &deck.Owner)
	return deck, err

}
