package handler

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

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", tasks)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if fmt.Sprintf("%d", task.ID) == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func Add(c *gin.Context) {
	text := c.PostForm("text")
	task := Task{ID: currentID, Text: text}
	currentID++
	tasks = append(tasks, task)
	c.Redirect(http.StatusSeeOther, "/")
}
