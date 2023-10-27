package booksapi

import (
	books "books/gen/books"
	context "context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type bookssrvc struct {
	logger *log.Logger
	db     *sql.DB
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	// Create a database connection.
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") // Replace with your database connection details
	if err != nil {
		log.Fatal(err)
	}
	return &bookssrvc{
		logger: logger,
		db:     db,
	}
}

// Close the database connection when the service is closed.
func (s *bookssrvc) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

// Create a reusable database query function.
func (s *bookssrvc) executeQuery(query string, args ...interface{}) error {
	_, err := s.db.Exec(query, args...)
	return err
}

// Create implements create.
func (s *bookssrvc) Create(ctx context.Context, p *books.Book) (res *books.Book, err error) {
	s.logger.Print("books.create")

	// Create the SQL query for inserting a new book
	insertQuery := "INSERT INTO books (Title, Author, BookCover, PublishedAt) VALUES (?, ?, ?, ?)"

	// Execute the query using the reusable function
	if err := s.executeQuery(insertQuery, *p.Title, *p.Author, *p.BookCover, *p.PublishedAt); err != nil {
		return nil, err
	}

	// Return the created book
	return p, nil
}

// All implements all.
func (s *bookssrvc) All(ctx context.Context) (res []*books.Book, err error) {
	s.logger.Print("books.all")

	// Create the SQL query for retrieving all books
	selectQuery := "SELECT * FROM books"

	// Execute the query and handle the results
	rows, err := s.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var booksList []*books.Book

	for rows.Next() {
		var book books.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.BookCover, &book.PublishedAt)
		if err != nil {
			return nil, err
		}
		booksList = append(booksList, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return booksList, nil
}

// UpdateBook implements updateBook.
func (s *bookssrvc) UpdateBook(ctx context.Context, p *books.UpdateBookPayload) (res *books.Book, err error) {
	s.logger.Print("books.updateBook")

	// Create the SQL query for updating a book
	updateQuery := "UPDATE books SET Title = ?, Author = ?, BookCover = ?, PublishedAt = ? WHERE Id = ?"

	// Execute the query using the reusable function
	if err := s.executeQuery(updateQuery, p.Book.Title, p.Book.Author, p.Book.BookCover, p.Book.PublishedAt, p.ID); err != nil {
		return nil, err
	}

	// Return the updated book
	return p.Book, nil
}

// GetBook implements getBook.
func (s *bookssrvc) GetBook(ctx context.Context, p *books.GetBookPayload) (res *books.Book, err error) {
	s.logger.Print("books.getBook")

	// Create the SQL query for retrieving a book by ID
	selectQuery := "SELECT * FROM books WHERE Id = ?"

	// Execute the query and handle the result
	var book books.Book
	err = s.db.QueryRow(selectQuery, p.ID).Scan(&book.ID, &book.Title, &book.Author, &book.BookCover, &book.PublishedAt)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// DeleteBook implements deleteBook.
func (s *bookssrvc) DeleteBook(ctx context.Context, p *books.DeleteBookPayload) (err error) {
	s.logger.Print("books.deleteBook")

	// Create the SQL query for deleting a book by ID
	deleteQuery := "DELETE FROM books WHERE Id = ?"

	// Execute the query using the reusable function
	if err := s.executeQuery(deleteQuery, p.ID); err != nil {
		return err
	}

	return nil
}
