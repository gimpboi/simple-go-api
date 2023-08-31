package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// film represents data about a film
type film struct {
     Title string `json:"title"`
	 Director string `json:"director"`
     Released string `json:"released"`
     Genre string `json:"genre"`
}

// test data
var films = []film{
    {Title: "Mulholland Drive",Director: "David Lynch",Released: "2001", Genre: "Detective Thriller"},
    {Title: "Schindler's List",Director: "Steven Spielberg",Released: "1993",Genre: "Holocaust film"},
    {Title: "Who Killed Paranoid Man?",Director: "Narty",Released: "900",Genre: "Tragedy"},
}

func main() {
	// Initialize a Gin router as Default
	router := gin.Default()
	// Use a GET function to associate the GET HTTP method and the /albums path with a handler function
	router.GET("/films",getFilms)
	// Use a POST function to associate the POST HTTP method and the /albums path with a handler function
	router.POST("/films",postFilms)
	// Use the RUN function to attach the router to an http.Server and start the server
	router.Run("localhost:8080")
}

// getFilms responds with a list of all films as indented JSON
func getFilms(c *gin.Context) {
     c.IndentedJSON(http.StatusOK, films)
}

// postFilms appends a film from JSON received in the request body
func postFilms(c *gin.Context) {
	var newFilm film

	// bind received JSON to newFilm
	if err := c.BindJSON(&newFilm); err != nil {
		return
	}

	// Add the new film to the slice
	films = append(films, newFilm)
	c.IndentedJSON(http.StatusCreated, newFilm)
}
