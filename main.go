package main
package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
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

var albums = []album{
    {ID: "1", Title: "Sempiternal", Artist:"Bring Me The Horizon", Price: 47.23},
    {ID: "2", Title: "Asking Alexandria", Artist: "Asking Alexandria", Price: 32.58},
    {ID: "3", Title: "10,000 Days", Artist: "TOOL", Price: 62.34},
}

var indexRoute string = "index"

func getIndex(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, indexRoute)
}


func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
	if a.ID == id {
	c.IndentedJSON(http.StatusOK, a)
	return
	}
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}
