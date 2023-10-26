package main

import (
	"fmt"
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
	router.DELETE("/clubs/:id", deleteClubById)
	router.GET("/health", health)

	router.Run("0.0.0.0:8080")
}

func getClubs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clubs)
}

func addClub(c *gin.Context) {
	var newClub club

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

func deleteClubById(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("delete for ", id, " clubs ", len(clubs))
	if i, err := strconv.Atoi(id); err == nil {

		clubs = append(clubs[:i-1], clubs[i:]...)
		fmt.Println("after deletion: ", clubs)
		c.Status(http.StatusOK)
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
}

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{ Status string }{Status: "OK"})

}
