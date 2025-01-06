package controllers

import (
	"context"
	"net/http"
	"time"

	"desafio-taghos-backend-jr/config"
	"desafio-taghos-backend-jr/database"
	"desafio-taghos-backend-jr/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection

func init() {
	getBooks := config.GetBooks()
	database.ConnectMongoDB()
	bookCollection = database.MongoClient.Database(getBooks).Collection(getBooks)	
}

// Create a book
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	book.ID = primitive.NewObjectID()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	_, err := bookCollection.InsertOne(context.TODO(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// Update a book
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	book.UpdatedAt = time.Now()

	_, err = bookCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": book},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// Get a book
func FindBook(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	err = bookCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// Get all books
func FindAllBooks(c *gin.Context) {
	cursor, err := bookCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer cursor.Close(context.TODO())

	var books []models.Book
	for cursor.Next(context.TODO()) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// Delete a book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = bookCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
