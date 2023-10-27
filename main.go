package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID   int
	Text string
}

var tasks []Task
var currentID = 1

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", tasks)
	})

	router.POST("/add", func(c *gin.Context) {
		text := c.PostForm("text")
		task := Task{ID: currentID, Text: text}
		currentID++
		tasks = append(tasks, task)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.GET("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, task := range tasks {
			if fmt.Sprintf("%d", task.ID) == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
