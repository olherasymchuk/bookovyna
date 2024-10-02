package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olherasymchuk/bookovyna/data"
)

func allBooks(c *gin.Context) {
	if result := data.Base.Find(&data.Books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &data.Books)
}

func oneBookByID(c *gin.Context) {
	if result := data.Base.First(&data.Books, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &data.Books)
}

func addBook(c *gin.Context) {
	var new data.Book
	c.Bind(&new)

	book := data.Book{Title: new.Title, Author_ID: new.Author_ID, Price: new.Price, Publisher_ID: new.Publisher_ID, Published: new.Published, ISBN: new.ISBN}
	result := data.Base.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Added", "post": book})
}
