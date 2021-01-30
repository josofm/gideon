package deck

import (
	"fmt"
	"log"

	"github.com/josofm/gideon/model"
	"gorm.io/gorm"
)

type DeckRepository struct{}

func (d *DeckRepository) Create(deck model.Deck, dbPool *gorm.DB) (string, error) {
	insertStatment := `INSERT INTO "deck" (name, owner, commander, cards) VALUES ($1, $2, $3, $4) RETURNING id;`
	err := dbPool.QueryRow(insertStatment, deck.Name, deck.Owner.ID, deck.Commander.Card.MultiverseId, deck.Cards[0].Card.MultiverseId).Scan(&deck.ID)
	if err != nil {
		fmt.Println(err)
		log.Print("[Create Deck] Fail database insertion")
		return "", err
	}
	return deck.Name, nil
}

// func (d *DeckRepository) UpdateDeck(db *gorm.DB, deck model.Deck) (int64, error) {
// 	result, err := db.Exec("update decks set owner=$1, commander=$2 where id=$3 RETURNING id",
// 		&deck.Owner, &deck.Commander, &deck.ID)

// 	if err != nil {
// 		return 0, err
// 	}

// 	rowsUpdated, err := result.RowsAffected()

// 	if err != nil {
// 		return 0, nil
// 	}

// 	return rowsUpdated, nil
// }

// func (d *DeckRepository) RemoveDeck(db *gorm.DB, id int) (int64, error) {
// 	result, err := db.Exec("delete from decks where id = $1", id)
// 	if err != nil {
// 		return 0, err
// 	}

// 	rowsDeleted, err := result.RowsAffected()

// 	if err != nil {
// 		return 0, nil
// 	}

// 	return rowsDeleted, nil

// }
