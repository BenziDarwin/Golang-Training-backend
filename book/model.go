package book

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{}
var nextID = 1

// GetBooksFromMemory returns the current list of books
func GetBooksFromMemory() []Book {
	return books
}
