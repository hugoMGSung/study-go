package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type hero struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	HeroName string  `json:"heroName"`
	Power    float64 `json:"power"`
}

var heros = []hero{
	{ID: "1", Name: "Tony Stark", HeroName: "IronMan", Power: 8000.0},
	{ID: "2", Name: "Steve Rogers", HeroName: "Captain America", Power: 6500.0},
	{ID: "3", Name: "Bruce Banner", HeroName: "Hulk", Power: 12000.0},
}

func getHeros(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, heros)
}

func postHeros(c *gin.Context) {
	var newHero hero

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newHero); err != nil {
		return
	}

	heros = append(heros, newHero)
	c.IndentedJSON(http.StatusCreated, newHero)
}

func main89() {
	router := gin.Default()
	router.GET("/heros", getHeros)
	router.POST("/heros", postHeros)

	router.Run(":8889")
}
