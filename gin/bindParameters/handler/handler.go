package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AssignRoute(e *gin.Engine) {
	g := e.Group("hoge")
	{
		g.GET("/array/param", h.GetArrayParamPath)
		g.GET("/param/:id", h.GetParamUriPath)
	}
}
