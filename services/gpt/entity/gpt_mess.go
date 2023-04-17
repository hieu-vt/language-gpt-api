package entity

import (
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
)

type GptContexts struct {
	core.SQLModel
	UserId      int    `json:"-" gorm:"column:user_id"`
	Uid         string `json:"userId" gorm:"-"`
	SendMessage string `json:"sendMessage" gorm:"column:send_message"`
	GptMessage  string `json:"gptMessage" gorm:"column:gpt_message"`
}

func (GptContexts) TableName() string {
	return "contexts"
}

func (g *GptContexts) MaskGpt() {
	uid := core.NewUID(uint32(g.UserId), common.MaskTypeUser, 1)
	g.Uid = uid.String()
	g.Mask(common.MaskTypeGpt)
}

type CreateGptContexts struct {
	SendMessage string `json:"sendMessage" gorm:"column:send_message"`
	GptMessage  string `json:"gptMessage" gorm:"column:gpt_message"`
}
