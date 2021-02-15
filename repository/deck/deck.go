package deck

import (
	"log"

	"github.com/josofm/gideon/model"
	"gorm.io/gorm"
)

type DeckRepository struct{}

func (d *DeckRepository) Create(deck model.Deck, dbPool *gorm.DB) (string, error) {
	if result := dbPool.Create(&deck); result.Error != nil {
		log.Print("[Create Deck] Fail database insertion")
		return "", result.Error
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
