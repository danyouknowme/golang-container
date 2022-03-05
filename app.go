package main

import (
	"net/http"
	"os"

	"golang-container-demo/myMusic"

	"github.com/gin-gonic/gin"
)

// getMySongs responds with the list of all songs as JSON.
func getMySongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, myMusic.GetMySongs())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Show index page
func showIndexPage(c *gin.Context) {
	env := os.Getenv("NODE_ENV")
	dat, err := os.ReadFile("./data.txt")
	check(err)

	html := `<div style="position: fixed; width: 100%; height: 100%; display: flex; justify-content: center; align-items: center; flex-direction: column;">
		<h1>Hello world, everyone!!</h1>
		<h3>from Golang ;)</h3>
		<hr/>
		<h4>NODE_ENV:
	`
	html = html + env
	html = html + "</h4><h4>data.txt: "
	html = html + string(dat)
	html = html + `</h4></div>`

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

// main
func main() {
	port := os.Getenv("PORT")

	router := gin.Default()
	router.GET("/", showIndexPage)
	router.GET("/my-songs", getMySongs)

	router.Run("0.0.0.0:" + port)
}
