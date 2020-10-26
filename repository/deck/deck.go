package deck

import (
	"database/sql"

	"github.com/josofm/gideon/model"
)

type DeckRepository struct{}

func NewDeckRepository() *DeckRepository {
	return &DeckRepository{}
}

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

	if err != nil {
		return model.Deck{}, err
	}

	return deck, nil

}

func (d DeckRepository) AddDeck(db *sql.DB, deck model.Deck) (int, error) {
	err := db.QueryRow("insert into decks (commander, owner) values ($1, $2) RETURNING id;",
		deck.Commander, deck.Owner).Scan(&deck.ID)

	if err != nil {
		return 0, err
	}

	return deck.ID, nil
}

func (d DeckRepository) UpdateDeck(db *sql.DB, deck model.Deck) (int64, error) {
	result, err := db.Exec("update decks set owner=$1, commander=$2 where id=$3 RETURNING id",
		&deck.Owner, &deck.Commander, &deck.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, nil
	}

	return rowsUpdated, nil
}

func (d *DeckRepository) RemoveDeck(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from decks where id = $1", id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, nil
	}

	return rowsDeleted, nil

}
