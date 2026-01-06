package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}

type SuccessResponse struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

type album struct {
	ID     string `json:"id"`
	Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// this is a slice here, this is so cool how the album data struct is populated
var albums = [] album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context){
	// the context is very important to carry request details
	c.IndentedJSON(http.StatusOK, albums) // good for serialisation
}

func postAlbums(c *gin.Context){
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil { // if the error is nil everything is fine
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	for _, val := range albums {
		if id == val.ID {
			c.IndentedJSON(http.StatusOK, SuccessResponse{
				Success: true,
				Data: val,
			})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Error:"album not found",
	})
}