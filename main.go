package main

import (
	"apis/book"
	"apis/student"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Student routes
	studentGroup := r.Group("/students")
	{
		studentGroup.POST("/", student.CreateStudent)
		studentGroup.GET("/", student.GetStudents)
		studentGroup.GET("/:id", student.GetStudentByID)
		studentGroup.PUT("/:id", student.UpdateStudent)
		studentGroup.DELETE("/:id", student.DeleteStudent)
		studentGroup.POST("/:id/borrow", student.BorrowBook)
	}

	// Book routes
	bookGroup := r.Group("/books")
	{
		bookGroup.POST("/", book.CreateBook)
		bookGroup.GET("/", book.GetBooks)
		bookGroup.GET("/:id", book.GetBookByID)
		bookGroup.PUT("/:id", book.UpdateBook)
		bookGroup.DELETE("/:id", book.DeleteBook)
	}

	r.Run(":8087")
}
