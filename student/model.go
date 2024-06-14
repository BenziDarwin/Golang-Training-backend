package student

import (
	"apis/book"
)

type Student struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Borrowed []book.Book `json:"borrowed"`
}

var students = []Student{}
var nextID = 1
