package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olherasymchuk/bookovyna/data"
)

func allAuthors(c *gin.Context) {
	if result := data.Base.Find(&data.Authors); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &data.Authors)
}

func oneAuthorByID(c *gin.Context) {
	if result := data.Base.First(&data.Authors, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &data.Authors)
}

func addAuthor(c *gin.Context) {
	var new data.Author
	c.Bind(&new)

	// Get data from request and create a new record
	author := data.Author{Name: new.Name, Surname: new.Surname}
	result := data.Base.Create(&author)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Added:", "post": author})
}
