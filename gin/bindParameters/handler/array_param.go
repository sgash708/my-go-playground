package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetArrayParamPathRequest struct {
	Params []int `form:"params[]", url:"params[]"`
}

func (h *Handler) GetArrayParamPath(gc *gin.Context) {
	var req GetArrayParamPathRequest
	if err := gc.Bind(&req); err != nil {
		gc.JSON(http.StatusBadRequest, err)
	}
	gc.JSON(http.StatusOK, req.Params)
}
