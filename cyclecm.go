package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gschei/cyclecm/domain/club"
	clubRepo "github.com/gschei/cyclecm/domain/club/memory"
)

var clubRepository club.ClubRepository = clubRepo.New()

func main() {
	club, _ := club.NewClub("Gilbert")
	clubRepository.Add(club)

	router := gin.Default()
	router.GET("/clubs", getClubs)
	router.POST("/clubs", addClub)
	router.GET("/clubs/:id", getClubById)
	//router.DELETE("/clubs/:id", deleteClubById)
	router.GET("/health", health)

	router.Run("0.0.0.0:8080")
}

func getClubs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clubRepository.GetAll())
}

func addClub(c *gin.Context) {
	var newClub club.Club

	if err := c.BindJSON(&newClub); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	newClub, err := clubRepository.Add(newClub)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newClub)
}

func getClubById(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	club, err := clubRepository.Get(int64(i))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, club)
}

/*
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
}*/

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{ Status string }{Status: "OK"})

}
