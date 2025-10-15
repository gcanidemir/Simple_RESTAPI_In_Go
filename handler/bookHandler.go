package handler

import (
	"Simple_RESTAPI_In_Go/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entity.Books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range entity.Books {
		if id == strconv.Itoa(book.ID) {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func PostBook(c *gin.Context) {
	var newBook entity.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	entity.Books = append(entity.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func UpdateBookByID(c *gin.Context) {
	id := c.Param("id")
	var updatedBook entity.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	for i, a := range entity.Books {
		if id == strconv.Itoa(a.ID) {
			entity.Books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func DeleteBookByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range entity.Books {
		if id == strconv.Itoa(a.ID) {
			entity.Books = append(entity.Books[:i], entity.Books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
