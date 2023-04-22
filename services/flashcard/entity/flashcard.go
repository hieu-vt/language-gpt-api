package entity

import (
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
)

type FlashCard struct {
	core.SQLModel
	UserId      int    `json:"-" gorm:"column:user_id"`
	Uid         string `json:"userId" gorm:"-"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (FlashCard) TableName() string {
	return "flashcards"
}

func (g *FlashCard) MaskFlash() {
	uid := core.NewUID(uint32(g.UserId), common.MaskTypeUser, 1)
	g.Uid = uid.String()
	g.Mask(common.MaskTypeFlashCard)
}

type CreateFlashCard struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}
