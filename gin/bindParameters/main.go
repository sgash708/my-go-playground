package main

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	e := gin.Default()

	h := handler.NewHandler()
	h.AssignRoute(e)
}
