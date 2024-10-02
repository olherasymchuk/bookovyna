package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olherasymchuk/bookovyna/data"
)

func allPublishers(c *gin.Context) {
	if result := data.Base.Find(&data.Publishers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &data.Publishers)
}

func onePublisherByID(c *gin.Context) {
	if result := data.Base.First(&data.Publishers, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &data.Publishers)
}

func addPublisher(c *gin.Context) {
	var new data.Publisher
	c.Bind(&new)

	publisher := data.Publisher{Name: new.Name}
	result := data.Base.Create(&publisher)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Added:", "post": publisher})
}
