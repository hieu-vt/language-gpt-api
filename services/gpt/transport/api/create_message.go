package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
	"lang-gpt-api/plugin/gpt"
	"lang-gpt-api/services/gpt/entity"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
	Save    bool   `json:"isSave"`
}

func (api *api) CreateMessage() func(c *gin.Context) {
	return func(c *gin.Context) {
		var body RequestBody
		_ = c.ShouldBind(&body)

		clientGpt := api.serviceCtx.MustGet(common.KeyCompGpt).(gpt.GptClient)

		err, dataGpt := clientGpt.RequestGptAPI(c, body.Message)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		if body.Save {
			go func(sendMessage string, gptMessage string) {
				_ = api.business.CreateMessageBiz(c, entity.CreateGptContexts{
					SendMessage: sendMessage,
					GptMessage:  gptMessage,
				})
			}(body.Message, dataGpt.Message)
		}

		c.JSON(http.StatusOK, core.ResponseData(entity.CreateGptContexts{
			SendMessage: body.Message,
			GptMessage:  dataGpt.Message,
		}))
	}
}
