package entity

import (
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
)

type Card struct {
	core.SQLModel
	UserId      int    `json:"-" gorm:"column:user_id"`
	Uid         string `json:"userId" gorm:"-"`
	FlashCardId int    `json:"-" gorm:"column:flashcard_id"`
	FCid        string `json:"flashCardId" gorm:"-"`
	FrontText   string `json:"frontText" gorm:"column:front_text"`
	BackText    string `json:"backText" gorm:"column:back_text"`
	Synonyms    string `json:"synonyms" gorm:"column:synonyms"`
}

func (Card) TableName() string {
	return "cards"
}

func (g *Card) MaskCard() {
	uid := core.NewUID(uint32(g.UserId), common.MaskTypeUser, 1)
	g.Uid = uid.String()
	cardId := core.NewUID(uint32(g.FlashCardId), common.MaskTypeFlashCard, 1)
	g.FCid = cardId.String()
	g.Mask(common.MaskTypeCard)
}

type CreateCard struct {
	FrontText string `json:"frontText" gorm:"column:front_text"`
	BackText  string `json:"backText" gorm:"column:back_text"`
	Synonyms  string `json:"synonyms" gorm:"column:synonyms"`
}
