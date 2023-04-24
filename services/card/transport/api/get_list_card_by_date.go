package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
	"net/http"
	"time"
)

func (a *api) GetListCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		dateString := c.Query("date")

		createdAt, err := time.Parse("02-01-2006", dateString)

		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := core.GetRequester(c)

		uid, _ := core.FromBase58(requester.GetSubject())
		requesterId := int(uid.GetLocalID()) // task owner id, id of who creates this new task

		err, result := a.business.GetCardByDate(c, requesterId, createdAt.Format("2006-01-02"))

		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		for i := range result {
			result[i].MaskCard()
		}

		c.JSON(http.StatusOK, core.ResponseData(result))
	}
}
