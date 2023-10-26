package booksapi

import (
	books "books/gen/books"
	"context"
	"log"
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
	res = &books.Book{}
	s.logger.Print("books.create")
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
