package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type club struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var clubs = []club{
	{ID: "1", Name: "Fleissige Radfahrer"},
	{ID: "2", Name: "Puppyton"},
}

func main() {
	router := gin.Default()
	router.GET("/clubs", getClubs)
	router.POST("/clubs", addClub)
	router.GET("/clubs/:id", getClubById)

	router.Run("localhost:8080")
}

func getClubs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clubs)
}

func addClub(c *gin.Context) {
	var newClub club

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newClub); err != nil {
		return
	}

	// Add the new album to the slice.
	clubs = append(clubs, newClub)
	c.IndentedJSON(http.StatusCreated, newClub)
}

func getClubById(c *gin.Context) {
	id := c.Param("id")

	if i, err := strconv.Atoi(id); err == nil {
		c.IndentedJSON(http.StatusOK, clubs[i-1])
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

}
