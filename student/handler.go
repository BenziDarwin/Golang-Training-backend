package student

import (
	"apis/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStudent handles the creation of a new student
func CreateStudent(c *gin.Context) {
	var newStudent Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStudent.ID = nextID
	nextID++
	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

// GetStudents retrieves all students
func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

// GetStudentByID retrieves a student by ID
func GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// UpdateStudent updates a student by ID
func UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	var updatedStudent Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, student := range students {
		if student.ID == id {
			updatedStudent.ID = student.ID
			students[i] = updatedStudent
			c.JSON(http.StatusOK, updatedStudent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// DeleteStudent deletes a student by ID
func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// BorrowBook handles a student borrowing a book
func BorrowBook(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	var bookToBorrow book.Book
	if err := c.ShouldBindJSON(&bookToBorrow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, student := range students {
		if student.ID == studentID {
			for _, b := range book.GetBooksFromMemory() {
				if b.ID == bookToBorrow.ID {
					students[i].Borrowed = append(students[i].Borrowed, b)
					c.JSON(http.StatusOK, students[i])
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}
