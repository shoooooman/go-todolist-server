package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// Todo is one item in todo list
type Todo struct {
	gorm.Model
	Deadline string `json:"deadline"`
	Title    string `json:"title"`
	Memo     string `json:"memo"`
}

func getList(c *gin.Context) {
	c.JSON(200, gin.H{
		"events": dbGetAll(),
	})
}

func getTodo(c *gin.Context) {
	idStr := c.Param("id")
	want, _ := strconv.Atoi(idStr)
	todo, notFound := dbGetOne(want)
	if notFound {
		c.String(404, "Not Found\n")
	} else {
		c.JSON(200, gin.H{
			"id":       want,
			"deadline": todo.Deadline,
			"title":    todo.Title,
			"memo":     todo.Memo,
		})
	}
}

func addTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	if _, err := time.Parse(time.RFC3339, todo.Deadline); err != nil {
		c.JSON(400, gin.H{
			"status":  "failure",
			"message": "invalid data format",
		})
	} else {
		id := dbInsert(todo.Deadline, todo.Title, todo.Memo)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "registered",
			"id":      id,
		})
	}
}

func deleteList(c *gin.Context) {
	dbDeleteAll()
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "deleted",
	})
}

func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	want, _ := strconv.Atoi(idStr)
	_, notFound := dbGetOne(want)
	if notFound {
		c.String(404, "Not Found\n")
	} else {
		dbDeleteOne(want)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "deleted",
			"id":      want,
		})
	}
}

func dbSetUp() {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_todo")
	if err != nil {
		fmt.Println(err.Error())
		panic("Error in dbSetUp()")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

func dbInsert(deadline string, title string, memo string) uint {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_todo?parseTime=true")
	if err != nil {
		panic("Error in dbInsert()")
	}
	todo := Todo{Deadline: deadline, Title: title, Memo: memo}
	db.Create(&todo)
	defer db.Close()
	return todo.ID
}

func dbGetOne(id int) (Todo, bool) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_todo?parseTime=true")
	if err != nil {
		panic("Error in dbGetOne()")
	}
	var todo Todo
	notFound := db.First(&todo, id).RecordNotFound()
	defer db.Close()
	return todo, notFound
}

func dbGetAll() []Todo {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_todo?parseTime=true")
	if err != nil {
		panic("Error in dbGetAll()")
	}
	var list []Todo
	db.Find(&list)
	defer db.Close()
	return list
}

func dbDeleteAll() {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_todo?parseTime=true")
	if err != nil {
		panic("Error in dbDeleteAll()")
	}
	var todo []Todo
	db.Find(&todo)
	db.Delete(&todo)
	defer db.Close()
}

func dbDeleteOne(id int) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_todo?parseTime=true")
	if err != nil {
		panic("Error in dbDeleteOne()")
	}
	var todo Todo
	db.Find(&todo, id)
	db.Delete(&todo)
	defer db.Close()
}

func main() {
	dbSetUp()
	r := gin.Default()
	r.GET("/api/v1/event", getList)
	r.GET("/api/v1/event/:id", getTodo)
	r.POST("/api/v1/event", addTodo)
	r.DELETE("api/v1/event", deleteList)
	r.DELETE("/api/v1/event/:id", deleteTodo)
	r.Run()
}
