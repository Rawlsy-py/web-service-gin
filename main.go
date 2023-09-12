// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

type Todo struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func main() {
	// Connect to the database
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=your_user dbname=your_db password=your_password sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Todo{})

	r := gin.Default()

	r.GET("/todos", GetTodos)
	r.POST("/todos", CreateTodo)
	r.PUT("/todos/:id", UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func GetTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)

	db.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&todo)
	db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	db.Where("id = ?", id).Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
