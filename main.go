package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// film represents data about a film
type film struct {
     title string "json:title"
     director string "json:director"
     released string "json:released"
     genre string "json:genre"
}

// test data
var films = []film{
    {title: "Mulholland Drive",director: "David Lynch",released: "2001", genre: "Detective Thriller"},
    {title: "Schindler's List",director: "Steven Spielberg",released: "1993",genre: "Holocaust film"},
    {title: "Who Killed Paranoid Man?",director: "Narty",released: "900",genre: "Tragedy"},
}

func main() {
	// Initialize a Gin router as Default
	router := gin.Default()
	// Use a GET function to associate the GET HTTP method and the /albums path with a handler function
	router.GET("/films",getFilms)
	// Use the RUN function to attach the router to an http.Server and start the server
	router.Run("localhost:8080")
}

// getFilms responds with a list of all films as indented JSON
func getFilms(c *gin.Context) {
     c.IndentedJSON(http.StatusOK, films)
}
