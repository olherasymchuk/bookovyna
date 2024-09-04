package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type author struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type publisher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type book struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Author_ID    string  `json:"author_id"`
	Price        float64 `json:"price"`
	Publisher_ID string  `json:"publisher"`
	Published    int     `json:"published"`
	ISBN         string  `json:"isbn"`
}

// authors slice
var authors = []author{
	{
		ID:      "1",
		Name:    "Андрій",
		Surname: "Андрощук",
	},
	{
		ID:      "2",
		Name:    "Богдан",
		Surname: "Боровик",
	},
	{
		ID:      "3",
		Name:    "Вікторія",
		Surname: "Василенко",
	},
}

var publishers = []publisher{
	{
		ID:   "1",
		Name: "Анархія-друк",
	},
	{
		ID:   "2",
		Name: "Барка і штиль",
	},
	{
		ID:   "3",
		Name: "Вокабюляри нової доби",
	},
}

var books = []book{
	{
		ID:           "1",
		Title:        "Асканія-Нова. Історія заповідника",
		Author_ID:    "1",
		Price:        380.00,
		Publisher_ID: "1",
		Published:    2024,
		ISBN:         "1111111111111",
	},

	{
		ID:           "2",
		Title:        "Брати і кузени",
		Author_ID:    "2",
		Price:        300.00,
		Publisher_ID: "2",
		Published:    2024,
		ISBN:         "1111111111112",
	},
	{
		ID:           "3",
		Title:        "Віра і майна",
		Author_ID:    "3",
		Price:        350.00,
		Publisher_ID: "3",
		Published:    2024,
		ISBN:         "1111111111113",
	},
}

func main() {
	router := gin.Default()
	api_1 := router.Group("/api/1")

	api_1.GET("/authors", getauthors)
	api_1.GET("/authors/:id", getauthorByID)
	api_1.POST("/authors", postauthors)

	api_1.GET("/publishers", getpublishers)
	api_1.GET("/publishers/:id", getpublisherByID)
	api_1.POST("/publishers", postpublishers)

	api_1.GET("/books", getbooks)
	api_1.GET("/books/:id", getbookByID)
	api_1.POST("/books", postbooks)

	router.Run("localhost:8080")
}

func getauthors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, authors)
}

func getpublishers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, publishers)
}

func getbooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getauthorByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range authors {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Автор не знайдений"})
}

func getpublisherByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range publishers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Автор не знайдений"})
}

func getbookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не знайдена"})
}

func postauthors(c *gin.Context) {
	var newauthor author

	// Call BindJSON to bind the received JSON to
	// newauthor.
	if err := c.BindJSON(&newauthor); err != nil {
		return
	}

	// Add the new author to the slice.
	authors = append(authors, newauthor)
	c.IndentedJSON(http.StatusCreated, newauthor)
}

func postpublishers(c *gin.Context) {
	var newpublisher publisher

	if err := c.BindJSON(&newpublisher); err != nil {
		return
	}

	publishers = append(publishers, newpublisher)
	c.IndentedJSON(http.StatusCreated, newpublisher)
}

func postbooks(c *gin.Context) {
	var newbook book

	if err := c.BindJSON(&newbook); err != nil {
		return
	}

	books = append(books, newbook)
	c.IndentedJSON(http.StatusCreated, newbook)
}
