package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/olherasymchuk/bookovyna/data"
	"github.com/olherasymchuk/bookovyna/handlers"
	"github.com/stretchr/testify/assert"
)

var api_ver string = "/api/1"

// TestMain sets up and tears down resources for all tests.
func TestMain(m *testing.M) {
	data.ConnectDB()

	// Run all tests

	code := m.Run()

	// Additional teardown code here

	// Exit with the test result code
	os.Exit(code)
}
func TestPostAuthors(t *testing.T) {
	author := data.Author{
		Name:    "Григорій",
		Surname: "Горинич",
	}
	jsonValue, _ := json.Marshal(author)
	req, _ := http.NewRequest("POST", api_ver+"/authors", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestGetAuthors(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/authors", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, data.Authors)
}
func TestGetAuthorByID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/authors/1", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, data.Authors)
}
func TestGetAuthorByNotExistingID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/authors/404", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestPostPublishers(t *testing.T) {
	publisher := data.Publisher{
		Name: "Галичанська книга",
	}
	jsonValue, _ := json.Marshal(publisher)
	req, _ := http.NewRequest("POST", api_ver+"/publishers", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestGetPublishers(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/publishers", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, data.Publishers)
}
func TestGetPublisherByID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/publishers/1", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, data.Publishers)
}

func TestGetPublisherByNotExistingID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/publishers/404", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestPostBooks(t *testing.T) {
	book := data.Book{
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
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestGetBooks(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/books", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, data.Books)
}
func TestGetBookByID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/books/1", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, data.Books)
}
func TestGetBookByNotExistingID(t *testing.T) {
	req, _ := http.NewRequest("GET", api_ver+"/books/404", nil)
	w := httptest.NewRecorder()
	r := handlers.SetupRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
