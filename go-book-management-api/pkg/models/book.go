package models

import (
	"github.com/festus/go-bookmanagement-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate()
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func DeleteBook(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return &book
}
