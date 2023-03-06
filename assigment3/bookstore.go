package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Book{})

	router := gin.Default()

	router.GET("/books", getBooks)
	//curl -X GET http://localhost:8080/books
	router.GET("/books/:id", getBookByID)
	//curl -X GET http://localhost:8080/books/1
	router.PUT("/books/:id", updateBookByID)
	//curl -X PUT -H "Content-Type: application/json" -d '{"title":"New title", "description":"New description", "price": 100.0}' http://localhost:8080/books/1
	router.DELETE("/books/:id", deleteBookByID)
	//curl -X DELETE http://localhost:8080/books/1
	router.GET("/search", searchBooks)
	//curl -X GET http://localhost:8080/search?title=book
	router.POST("/books", addBook)
	//curl -X POST -H "Content-Type: application/json" -d '{"title":"New title", "description":"New description", "price": 100.0}' http://localhost:8080/books
	//curl -X POST -H "Content-Type: application/json" -d '[{"title":"New title", "description":"New description", "price": 100.0}, {"title":"New title", "description":"New description", "price": 100.0}]' http://localhost:8080/books
	//curl -X POST -H "Content-Type: application/json" -d '{"title":"New title", "description":"New description", "price": 100.0}' http://localhost:8080/books
	router.GET("/sorted-books/:order", getSortedBooks)
	//curl -X GET http://localhost:8080/sorted-books/asc

	router.Run(":8080")
}

func getBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book
	db.First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func updateBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book
	db.First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

func deleteBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book
	db.First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func searchBooks(c *gin.Context) {
	title := c.Query("title")
	var books []Book
	db.Where("title LIKE ?", "%"+title+"%").Find(&books)
	c.JSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusOK, book)
}

func getSortedBooks(c *gin.Context) {
	order := c.Param("order")
	var books []Book
	switch order {
	case "desc":
		db.Order("price DESC").Find(&books)
	case "asc":
		db.Order("price ASC").Find(&books)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort order"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Dockerfile
//docker build -t bookstore .
//docker run -p 8080:8080 bookstore
