package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	e := gin.Default()

	assignHandler(e)

	if err := e.Run(":80"); err != nil {
		panic(err)
	}
}

func assignHandler(e *gin.Engine) {
	h := e.Group("hoge")
	{
		h.GET("/path", GetHogePath)
	}
}

type GetHogePathRequest struct {
	Params []int `form:"params[]" url:"params[]"`
}

func GetHogePath(gc *gin.Context) {
	var req GetHogePathRequest
	if err := gc.Bind(&req); err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}
	gc.JSON(http.StatusOK, req.Params)
}
