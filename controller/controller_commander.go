package controller

import (
	"fmt"
	"strings"

	"github.com/josofm/mtg-sdk-go"
)

func isBanned(card *mtg.Card, format string) bool {
	for _, l := range card.Legalities {
		if l.Format == format && l.Legality == "Legal" {
			return false
		}
	}
	return true
}

func isValidCommander(multiverseID mtg.MultiverseId) bool {
	card, err := mtg.MultiverseId(multiverseID).Fetch()
	fmt.Println(err)
	fmt.Println(card)
	if err != nil {
		return false
	}
	if isBanned(card, "Commander") {
		return false
	}
	rarity := false
	if strings.Contains(card.Rarity, "Mythic") || strings.Contains(card.Rarity, "Rare") {
		rarity = true
	}
	if !rarity {
		return false
	}
	if strings.Contains(card.OriginalText, "can be your commander") {
		return true
	}
	legendary := false
	for _, t := range card.Types {
		if strings.Contains(t, "Legendary") {
			legendary = true
			break
		}
	}
	if !legendary {
		return false
	}
	return strings.Contains(card.Type, "Creature")
}
