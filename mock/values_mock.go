package mock

import (
	"github.com/josofm/gideon/model"

	"github.com/josofm/mtg-sdk-go"
)

func GetBasicDeck() model.Deck {
	return model.Deck{
		Name: "best deck",
		Commander: model.Card{
			Card: mtg.Card{
				MultiverseId: "389712",
			},
		},
		Cards: []model.Card{
			{
				Card: mtg.Card{
					MultiverseId: "409574",
				},
			},
		},
	}
}
