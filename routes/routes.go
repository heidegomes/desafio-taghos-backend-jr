package routes

import (
	"desafio-taghos-backend-jr/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/books", controllers.CreateBook)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.GET("/books/:id", controllers.FindBook)
    r.GET("/books", controllers.FindAllBooks)
    r.DELETE("/books/:id", controllers.DeleteBook)
    return r
}
