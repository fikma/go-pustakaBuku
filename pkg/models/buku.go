package models

import (
	"github.com/fikma/go-pustakaBuku/pkg/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Nama      string `gorm:"" json:"nama"`
	Penulis   string `json:"penulis"`
	Publikasi string `json:"publikasi"`
}

func init() {
	configs.Connect()

	db = configs.GetDb()
	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) *Book {
	db.Create(&book)

	return book
}

func GetAllBooks() []Book {
	var books []Book

	db.Find(&books)

	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book

	rDb := db.Where("ID=?", Id).Find(&getBook)

	return &getBook, rDb
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)

	return book
}
