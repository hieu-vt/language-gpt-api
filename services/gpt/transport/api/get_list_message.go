package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/common"
	"net/http"
)

func (api *api) GetListMessage() func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging core.Paging
		if err := c.ShouldBind(&paging); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		paging.Process()

		err, result := api.business.GetMessagesBiz(c, &paging)

		if err != nil {
			common.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		for i := range result {
			result[i].MaskGpt()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(result, paging, nil))
	}
}
