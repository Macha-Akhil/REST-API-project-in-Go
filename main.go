package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "Go Basics", Author: "John"},
	{ID: 2, Title: "Gin in Action", Author: "Jane"},
}

func main() {
	server := gin.Default()

	server.GET("/books", getBooks)
	server.GET("/books/:id", getBookByID)
	server.POST("/books", addBook)
	server.PUT("/books/:id", updateBook)
	server.DELETE("/books/:id", deleteBook)

	server.Run(":8080")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)

}

func getBookByID(c *gin.Context) {
	fmt.Println(c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newBook.ID = books[len(books)-1].ID + 1
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updated Book
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	for i, b := range books {
		if b.ID == id {
			books[i].Title = updated.Title
			books[i].Author = updated.Author
			c.JSON(http.StatusOK, books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
