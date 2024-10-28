package main

import (
	"errors"

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

func createBook(c *gin.Context) {
	// the context retains all details belonging to a request
	// so if you had query parameters, data payload and stuff it will keep them there
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		// here we tried to bind the json from the request data
		// to the new book by passing its pointer. if we got an error
		// we simply return
		return
	}

	books = append(books, newBook)              // if we didnt get an error and we succesfully bounded the json we append the new book
	c.IndentedJSON(http.StatusCreated, newBook) // return the newly created book with StatusCreated status code
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id") // "/books/2" -> 2 is the id here
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."}) // gin.H is a shortcut to allow us to write our own custom json. it maps to a type string to interface
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	// gin router setup
	router := gin.Default()
	router.GET("/books", getBooks) // if you make a GET request to /books then we call our function. POST requests for example will not trigger this
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById) // path parameter :param
	router.Run("localhost:8080")
}
