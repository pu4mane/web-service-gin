package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordServer struct {
	*gin.Engine
	albums Albums
}

func NewRecordServer() *RecordServer {
	r := new(RecordServer)

	router := gin.Default()

	router.GET("/albums", r.getAlbums)
	router.GET("/albums/:id", r.getAlbumByID)
	router.POST("/albums", r.postAlbums)

	r.Engine = router

	return r
}

// getAlbums responds with the list of all albums as JSON.
func (r *RecordServer) getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, r.albums)
}

// getAlbumByID locates the album whose ID value matches the id.
func (r *RecordServer) getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range r.albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func (r *RecordServer) postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	r.albums = append(r.albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
