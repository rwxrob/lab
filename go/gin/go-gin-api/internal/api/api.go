package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Albums slice to seed record album data.
var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan",
		Price: 39.99},
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var a Album

	if err := c.BindJSON(&a); err != nil {
		return
	}

	Albums = append(Albums, a) // would normally persist as well

	c.IndentedJSON(http.StatusCreated, a)

}

// PetAlbumByID locates the album whose ID value matches the argument
// passed to the id parameter sent by the client, then returns that
// album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Serve starts the API server.
func Serve(args ...string) {
	addr := "localhost:8080"
	if len(args) > 1 {
		addr = args[1]
	}
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", PostAlbums)
	router.Run(addr)
}
