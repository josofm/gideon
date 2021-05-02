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
