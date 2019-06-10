package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Todo is one item in todo list
type Todo struct {
	ID       int    `json:"id"`
	Deadline string `json:"deadline"`
	Title    string `json:"title"`
	Memo     string `json:"memo"`
}

var list []Todo
var id = 0

func getList(c *gin.Context) {
	c.JSON(200, gin.H{
		"events": list,
	})
}

func getTodo(c *gin.Context) {
	idStr := c.Param("id")
	want, _ := strconv.Atoi(idStr)
	if want >= id || want < 0 {
		c.String(404, "Not Found\n")
	} else {
		c.JSON(200, gin.H{
			"id":       list[want].ID,
			"deadline": list[want].Deadline,
			"title":    list[want].Title,
			"memo":     list[want].Memo,
		})
	}
}

func appendTodo(id int, deadline string, title string, memo string) {
	todo := Todo{id, deadline, title, memo}
	list = append(list, todo)
}

func addTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	if _, err := time.Parse(time.RFC3339, todo.Deadline); err != nil {
		c.JSON(400, gin.H{
			"status":  "failure",
			"message": "invalid data ",
			"id":      id,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "registered",
			"id":      id,
		})
		appendTodo(id, todo.Deadline, todo.Title, todo.Memo)
		id++
	}
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/event", getList)
	r.GET("/api/v1/event/:id", getTodo)
	r.POST("/api/v1/event", addTodo)
	r.Run()
}
