package model

import (
	"time"

	"github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Commander Card   `json:commander,omitempty`
	Partner   Card   `json:partner,omitempty`
	Owner     User   `json:owner,omitempty`
	OwnerID   int
	Cards     []Card `gorm:"many2many:deck_cards;" json:"cards,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Card struct {
	gorm.Model
	ID        uint     `gorm:"primaryKey" json:"id,omitempty`
	Card      mtg.Card `json:"card,omitempty"`
	Price     Price    `json:"price,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Price struct {
	gorm.Model
	CardID    int
	Minimun   float64 `json:"minimun,omitempty"`
	Average   float64 `json:"average,omitempty"`
	Maximum   float64 `json:"maximun,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Sex       string `json:"sex,omitempty"`
	Age       string `json:"age,omitempty"`
	Token     string `json:"token,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}
