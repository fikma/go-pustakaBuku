package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fikma/go-pustakaBuku/pkg/models"
	"github.com/fikma/go-pustakaBuku/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(response http.ResponseWriter, request *http.Request) {
	createBook := &models.Book{}

	utils.ParseBody(request, createBook)
	bookData := models.CreateBook(createBook)
	result, _ := json.Marshal(bookData)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	response.Write(result)

}

func GetAllBooks(response http.ResponseWriter, request *http.Request) {
	newBook := models.GetAllBooks()

	res, _ := json.Marshal(newBook)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	response.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error ketika parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	result, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func UpdateBookById(response http.ResponseWriter, request *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(request, updateBook)

	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error ketika parsing")
	}

	bookDetails, db := models.GetBookById(id)

	if updateBook.Nama != "" {
		bookDetails.Nama = updateBook.Nama
	}

	if updateBook.Penulis != "" {
		bookDetails.Penulis = updateBook.Penulis
	}

	if updateBook.Publikasi != "" {
		bookDetails.Publikasi = updateBook.Publikasi
	}

	db.Save(*bookDetails)

	result, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func DeleteBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error ketika parsing")
	}

	book := models.DeleteBook(id)

	result, _ := json.Marshal(book)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}
