package main

import (
	//"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	/*
		we use uppercase for the fields its because we export them
		this means they are public and can be viewed outside the module we're in
		but in our json we want them to be lowercase so we add the json field manually
		if we don't make the first letter uppercase when we return a book we just get an empty json
	*/
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	// indented json is for formatting so it looks nice
	// http.StatusOK is the status code we're sending
	// books is what we send as the data
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	// gin router setup
	router := gin.Default()
	router.GET("/books", getBooks) // if you make a GET request to /books then we call our function. POST requests for example will not trigger this
	router.Run("localhost:8080")
}
