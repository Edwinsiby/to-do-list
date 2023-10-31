package main

import (
	"list/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", handler.Index)

	router.POST("/add", handler.Add)

	router.GET("/delete/:id", handler.Delete)

	router.Run(":8080")
}
