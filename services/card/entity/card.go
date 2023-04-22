package entity

import (
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
)

type FlashCard struct {
	core.SQLModel
	FlashCardId int    `json:"-" gorm:"column:flashcard_id"`
	FCid        string `json:"flashCardId" gorm:"-"`
	FrontText   string `json:"frontText" gorm:"column:front_text"`
	BackText    string `json:"backText" gorm:"column:back_text"`
	Synonyms    string `json:"synonyms" gorm:"column:synonyms"`
}

func (FlashCard) TableName() string {
	return "cards"
}

func (g *FlashCard) MaskCard() {
	uid := core.NewUID(uint32(g.FlashCardId), common.MaskTypeFlashCard, 1)
	g.FCid = uid.String()
	g.Mask(common.MaskTypeCard)
}

type CreateCard struct {
	FrontText string `json:"frontText" gorm:"column:front_text"`
	BackText  string `json:"backText" gorm:"column:back_text"`
	Synonyms  string `json:"synonyms" gorm:"column:synonyms"`
}
