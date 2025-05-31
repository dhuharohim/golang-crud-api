package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhuharohim/golang-crud-api/database"
	"github.com/dhuharohim/golang-crud-api/models"
	"github.com/dhuharohim/golang-crud-api/utils"
	"github.com/gorilla/mux"
)

// CreateBook handles creating a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := database.DB.Create(&book).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create book", err)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "Book created successfully", book)
}

// GetBooks handles retrieving all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch books", err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Books retrieved successfully", books)
}

// GetBook handles retrieving a single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid book ID", err)
		return
	}

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Book not found", err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Book retrieved successfully", book)
}

// UpdateBook handles updating a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid book ID", err)
		return
	}

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Book not found", err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := database.DB.Save(&book).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update book", err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Book updated successfully", book)
}

// DeleteBook handles deleting a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid book ID", err)
		return
	}

	if err := database.DB.Delete(&models.Book{}, id).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete book", err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Book deleted successfully", nil)
}
