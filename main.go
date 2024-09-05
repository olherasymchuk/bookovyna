package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Author struct {
	gorm.Model
	Name    string
	Surname string
}
type Publisher struct {
	gorm.Model
	Name string
}
type Book struct {
	gorm.Model
	Title        string
	Author_ID    uint32
	Price        float32
	Publisher_ID uint32
	Published    uint32
	ISBN         string
}

var authors = []Author{
	{
		Name:    "Андрій",
		Surname: "Андрощук",
	},
	{
		Name:    "Богдан",
		Surname: "Боровик",
	},
	{
		Name:    "Вікторія",
		Surname: "Василенко",
	},
}
var publishers = []Publisher{
	{
		Name: "Анархія-друк",
	},
	{
		Name: "Барка і штиль",
	},
	{
		Name: "Вокабюляри нової доби",
	},
}
var books = []Book{
	{
		Title:        "Асканія-Нова. Історія заповідника",
		Author_ID:    1,
		Price:        380.00,
		Publisher_ID: 1,
		Published:    2024,
		ISBN:         "1111111111111",
	},
	{
		Title:        "Брати і кузени",
		Author_ID:    2,
		Price:        300.00,
		Publisher_ID: 2,
		Published:    2024,
		ISBN:         "1111111111112",
	},
	{
		Title:        "Віра і майна",
		Author_ID:    3,
		Price:        350.00,
		Publisher_ID: 3,
		Published:    2024,
		ISBN:         "1111111111113",
	},
}

func main() {

	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Publisher{})
	db.AutoMigrate(&Book{})
	db.Create(authors)
	db.Create(publishers)
	db.Create(books)

	router := gin.Default()

	router.StaticFile("/favicon.ico", "./favicon.ico")

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
	if result := db.Find(&authors); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &authors)
}
func getpublishers(c *gin.Context) {
	if result := db.Find(&publishers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &publishers)
}
func getbooks(c *gin.Context) {
	if result := db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &books)
}
func getauthorByID(c *gin.Context) {
	if result := db.First(&authors, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &authors)
}
func getpublisherByID(c *gin.Context) {
	if result := db.First(&publishers, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &publishers)
}
func getbookByID(c *gin.Context) {
	if result := db.First(&books, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, &books)
}
func postauthors(c *gin.Context) {
	// Create a validating user's input:
	var new Author
	c.Bind(&new)

	// Get data from request and create a new record
	author := Author{Name: new.Name, Surname: new.Surname}
	result := db.Create(&author)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added:", "post": author})
}
func postpublishers(c *gin.Context) {

	var new Publisher
	c.Bind(&new)

	publisher := Publisher{Name: new.Name}
	result := db.Create(&publisher)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added:", "post": publisher})
}
func postbooks(c *gin.Context) {

	var new Book
	c.Bind(&new)

	book := Book{Title: new.Title, Author_ID: new.Author_ID, Price: new.Price, Publisher_ID: new.Publisher_ID, Published: new.Published, ISBN: new.ISBN}
	result := db.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added", "post": book})
}
