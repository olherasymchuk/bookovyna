package handlers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.StaticFile("/favicon.ico", "./favicon.ico")

	api_1 := router.Group("/api/1")

	api_1.GET("/authors", allAuthors)
	api_1.GET("/authors/:id", oneAuthorByID)
	api_1.POST("/authors", addAuthor)

	api_1.GET("/publishers", allPublishers)
	api_1.GET("/publishers/:id", onePublisherByID)
	api_1.POST("/publishers", addPublisher)

	api_1.GET("/books", allBooks)
	api_1.GET("/books/:id", oneBookByID)
	api_1.POST("/books", addBook)

	return router
}
