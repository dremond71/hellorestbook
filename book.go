package hellorestbook

import (
	"github.com/dremond71/hellorestdatabase"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := hellorestdatabase.DB
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := hellorestdatabase.DB
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := hellorestdatabase.DB
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := hellorestdatabase.DB

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
