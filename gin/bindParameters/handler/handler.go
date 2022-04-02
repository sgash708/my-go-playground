package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	// Application application.ApplicationInterface
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AssignRoute(e *gin.Engine) {
	g := e.Group("hoge")
	{
		g.GET("/array/param", GetArrayParamPath)
	}
}