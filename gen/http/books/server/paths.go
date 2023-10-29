// Code generated by goa v3.13.2, DO NOT EDIT.
//
// HTTP request path constructors for the books service.
//
// Command:
// $ goa gen books/design

package server

import (
	"fmt"
)

// CreateBooksPath returns the URL path to the books service create HTTP endpoint.
func CreateBooksPath() string {
	return "/books"
}

// AllBooksPath returns the URL path to the books service all HTTP endpoint.
func AllBooksPath() string {
	return "/books"
}

// UpdateBookBooksPath returns the URL path to the books service updateBook HTTP endpoint.
func UpdateBookBooksPath(id int) string {
	return fmt.Sprintf("/books/%v", id)
}

// GetBookBooksPath returns the URL path to the books service getBook HTTP endpoint.
func GetBookBooksPath(id int) string {
	return fmt.Sprintf("/books/%v", id)
}

// DeleteBookBooksPath returns the URL path to the books service deleteBook HTTP endpoint.
func DeleteBookBooksPath(id int) string {
	return fmt.Sprintf("/books/%v", id)
}

// UploadBooksPath returns the URL path to the books service upload HTTP endpoint.
func UploadBooksPath(dir string) string {
	return fmt.Sprintf("/upload/%v", dir)
}
