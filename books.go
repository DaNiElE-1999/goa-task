package booksapi

import (
	books "books/gen/books"
	"context"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// books service example implementation.
// The example methods log the requests and return zero values.
type bookssrvc struct {
	logger *log.Logger
}

func ExecuteQuery(query string) {
	godotenv.Load(".env")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") //toBeFixed with env var
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	return &bookssrvc{logger}
}

// Create implements create.
func (s *bookssrvc) Create(ctx context.Context, p *books.Book) (res *books.Book, err error) {
	res = &books.Book{}
	s.logger.Print("books.create")
	s.logger.Print("Received POST request with the following data:")
	s.logger.Printf("ID: %d", *p.ID)
	s.logger.Printf("Title: %s", *p.Title)
	s.logger.Printf("Author: %s", *p.Author)
	s.logger.Printf("BookCover: %s", *p.BookCover)
	s.logger.Printf("PublishedAt: %s", *p.PublishedAt)
	return
}

// All implements all.
func (s *bookssrvc) All(ctx context.Context) (res []*books.Book, err error) {
	s.logger.Print("books.all")
	return
}

// UpdateBook implements updateBook.
func (s *bookssrvc) UpdateBook(ctx context.Context, p *books.UpdateBookPayload) (res *books.Book, err error) {
	res = &books.Book{}
	s.logger.Print("books.updateBook")
	return
}

// GetBook implements getBook.
func (s *bookssrvc) GetBook(ctx context.Context, p *books.GetBookPayload) (res *books.Book, err error) {
	res = &books.Book{}
	s.logger.Print("books.getBook")
	return
}

// DeleteBook implements deleteBook.
func (s *bookssrvc) DeleteBook(ctx context.Context, p *books.DeleteBookPayload) (err error) {
	s.logger.Print("books.deleteBook")
	return
}
