package main

import (
	"database/sql"
	"web-service-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.CreateTables(db)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
