package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
	"lang-gpt-api/services/card/entity"
	"net/http"
)

func (a *api) CreateCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body entity.CreateCard

		if err := c.ShouldBind(&body); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := a.business.CreateCard(c, entity.Card{
			FrontText: body.FrontText,
			BackText:  body.BackText,
			Synonyms:  body.Synonyms,
		}); err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		c.JSON(http.StatusCreated, core.ResponseData(true))
	}
}
