package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    Price float64 `json:"price"`
}

func main() {
    router := gin.Default()
    router.GET("/", getIndex)
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

var toDo = []toDo{
}

var indexRoute string = "index"

func getIndex(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, indexRoute)
}

