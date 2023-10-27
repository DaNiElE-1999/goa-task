package booksapi

import (
	books "books/gen/books"
	"context"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// books service example implementation.
// The example methods log the requests and return zero values.
type bookssrvc struct {
	logger *log.Logger
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	return &bookssrvc{logger}
}

// Create implements create.
func (s *bookssrvc) Create(ctx context.Context, p *books.Book) (res *books.Book, err error) {
	s.logger.Print("books.create")

	// Execute the SQL query to insert a new book
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") // Replace with your database connection details
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Create the SQL query for inserting a new book
	insertQuery := "INSERT INTO books (Title, Author, BookCover, PublishedAt) VALUES (?, ?, ?, ?)"

	// Execute the query
	_, err = db.Exec(insertQuery, *p.Title, *p.Author, *p.BookCover, *p.PublishedAt)
	if err != nil {
		return nil, err
	}

	// Return the created book
	res = p
	return
}

// All implements all.
func (s *bookssrvc) All(ctx context.Context) (res []*books.Book, err error) {
	s.logger.Print("books.all")

	// Create a slice to store the results
	var booksList []*books.Book

	// Execute the SQL query to retrieve all books
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") // Replace with your database connection details
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Create the SQL query for retrieving all books
	selectQuery := "SELECT * FROM books"

	// Execute the query
	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and scan the data into the books slice
	for rows.Next() {
		var book books.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.BookCover, &book.PublishedAt)
		if err != nil {
			return nil, err
		}
		booksList = append(booksList, &book)
	}

	// Check for errors during row iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return booksList, nil
}

// UpdateBook implements updateBook.
func (s *bookssrvc) UpdateBook(ctx context.Context, p *books.UpdateBookPayload) (res *books.Book, err error) {
	s.logger.Print("books.updateBook")

	//Execute the SQL query to update the book
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") // Replace with your database connection details
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//Create the SQL query for updating a book
	updateQuery := "UPDATE books SET Title = ?, Author = ?, BookCover = ?, PublishedAt = ? WHERE Id = ?"

	//Execute the query
	_, err = db.Exec(updateQuery, p.Book.Title, p.Book.Author, p.Book.BookCover, p.Book.PublishedAt, p.ID)
	if err != nil {
		return nil, err
	}

	//Return the updated book
	return p.Book, nil
	//return
}

// GetBook implements getBook.
func (s *bookssrvc) GetBook(ctx context.Context, p *books.GetBookPayload) (res *books.Book, err error) {
	s.logger.Print("books.getBook")

	// Create a variable to store the retrieved book
	var book books.Book

	// Execute the SQL query to retrieve a book by ID
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") // Replace with your database connection details
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Create the SQL query for retrieving a book by ID
	selectQuery := "SELECT * FROM books WHERE Id = ?"

	// Execute the query
	err = db.QueryRow(selectQuery, p.ID).Scan(&book.ID, &book.Title, &book.Author, &book.BookCover, &book.PublishedAt)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// DeleteBook implements deleteBook.
func (s *bookssrvc) DeleteBook(ctx context.Context, p *books.DeleteBookPayload) (err error) {
	s.logger.Print("books.deleteBook")

	// Execute the SQL query to delete a book by ID
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") // Replace with your database connection details
	if err != nil {
		return err
	}
	defer db.Close()

	// Create the SQL query for deleting a book by ID
	deleteQuery := "DELETE FROM books WHERE Id = ?"

	// Execute the query
	_, err = db.Exec(deleteQuery, p.ID)
	if err != nil {
		return err
	}

	return nil
}
