package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var api_ver string = "/api/1"

// TestMain sets up and tears down resources for all tests.
func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Run all tests

	code := m.Run()

	// Additional teardown code here

	// Exit with the test result code
	os.Exit(code)
}
func TestPostAuthors(t *testing.T) {
	db.AutoMigrate(&Author{})
	db.Create(authors)

	author := Author{
		Name:    "Григорій",
		Surname: "Горинич",
	}
	jsonValue, _ := json.Marshal(author)
	req, _ := http.NewRequest("POST", api_ver+"/authors", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestGetAuthors(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/authors", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, authors)
}
func TestGetAuthorByID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/authors/1", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, authors)
}
func TestGetAuthorByNotExistingID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/authors/404", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestPostPublishers(t *testing.T) {
	db.AutoMigrate(&Publisher{})
	db.Create(publishers)

	publisher := Publisher{
		Name: "Галичанська книга",
	}
	jsonValue, _ := json.Marshal(publisher)
	req, _ := http.NewRequest("POST", api_ver+"/publishers", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestGetPublishers(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/publishers", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, publishers)
}
func TestGetPublisherByID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/publishers/1", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, publishers)
}

func TestGetPublisherByNotExistingID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/publishers/404", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestPostBooks(t *testing.T) {
	db.AutoMigrate(&Book{})
	db.Create(books)
	book := Book{
		Title:        "Асканія-Нова. Історія заповідника",
		Author_ID:    1,
		Price:        380.00,
		Publisher_ID: 1,
		Published:    2024,
		ISBN:         "1111111111111",
	}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", api_ver+"/books", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestGetBooks(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/books", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)
}
func TestGetBookByID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/books/1", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, books)
}
func TestGetBookByNotExistingID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/books/404", nil)
	w := httptest.NewRecorder()
	r := SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
