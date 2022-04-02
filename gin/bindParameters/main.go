package main

import (
	"bindParameters/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	e := gin.Default()

	h := handler.NewHandler()
	h.AssignRoute(e)

	if err := e.Run(":80"); err != nil {
		panic(err)
	}
}
