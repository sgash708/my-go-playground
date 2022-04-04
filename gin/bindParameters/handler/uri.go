package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetParamUriRequest struct {
	ID string `uri:"id"`
}

func (h *Handler) GetParamUriPath(gc *gin.Context) {
	var req GetParamUriRequest
	if err := gc.ShouldBindUri(&req); err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}
	gc.JSON(http.StatusOK, req.ID)
}
