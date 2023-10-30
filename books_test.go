package booksapi

import (
	"books/gen/books"
	"context"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type Book struct {
	Title       string
	Author      string
	BookCover   string
	PublishedAt string
}

var testBook = Book{
	Title:       "A Game of Thrones",
	Author:      "George R. R. Martin",
	BookCover:   "https://www.example.com/a_game_of_thrones.jpg",
	PublishedAt: "1996-08-01",
}

func TestCreateBook(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	// Create a new logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Create a new books service with the mock database connection and logger
	service := &bookssrvc{
		logger: logger,
		db:     db,
	}

	// Define the input and output for the test case
	input := &books.Book{
		Title:       &testBook.Title,
		Author:      &testBook.Author,
		BookCover:   &testBook.BookCover,
		PublishedAt: &testBook.PublishedAt,
	}
	expectedOutput := &books.Book{
		ID:          nil,
		Title:       &testBook.Title,
		Author:      &testBook.Author,
		BookCover:   &testBook.BookCover,
		PublishedAt: &testBook.PublishedAt,
	}

	// Set up the mock database to expect the insert query
	mock.ExpectExec("^INSERT INTO books").WithArgs(input.Title, input.Author, input.BookCover, input.PublishedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the Create method on the books service
	output, err := service.Create(context.Background(), input)

	// Check for errors and unexpected output
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if output.ID != expectedOutput.ID {
		t.Errorf("Unexpected ID: got %v, want %v", output.ID, expectedOutput.ID)
	}
	if output.Title != expectedOutput.Title {
		t.Errorf("Unexpected title: got %v, want %v", output.Title, expectedOutput.Title)
	}
	if output.Author != expectedOutput.Author {
		t.Errorf("Unexpected author: got %v, want %v", output.Author, expectedOutput.Author)
	}
	if output.BookCover != expectedOutput.BookCover {
		t.Errorf("Unexpected book cover: got %v, want %v", output.BookCover, expectedOutput.BookCover)
	}
	if output.PublishedAt != expectedOutput.PublishedAt {
		t.Errorf("Unexpected published date: got %v, want %v", output.PublishedAt, expectedOutput.PublishedAt)
	}

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestAllBooks(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	// Create a new logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Create a new books service with the mock database connection and logger
	service := &bookssrvc{
		logger: logger,
		db:     db,
	}

	// Define the input and output for the test case
	expectedOutput := []*books.Book{
		{
			ID:          nil,
			Title:       &testBook.Title,
			Author:      &testBook.Author,
			BookCover:   &testBook.BookCover,
			PublishedAt: &testBook.PublishedAt,
		},
		// Add more books as needed...
	}

	// Set up the mock database to expect the select query
	rows := sqlmock.NewRows([]string{"ID", "Title", "Author", "BookCover", "PublishedAt"}).
		AddRow(expectedOutput[0].ID, expectedOutput[0].Title, expectedOutput[0].Author, expectedOutput[0].BookCover, expectedOutput[0].PublishedAt)
	mock.ExpectQuery("^SELECT \\* FROM books").WillReturnRows(rows)

	// Call the All method on the books service
	output, err := service.All(context.Background())

	// Check for errors and unexpected output
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(output) != len(expectedOutput) {
		t.Errorf("Unexpected number of results: got %d, want %d", len(output), len(expectedOutput))
	}
	for i := range output {
		if output[i].ID != expectedOutput[i].ID {
			t.Errorf("Unexpected ID at index %d: got %v, want %v", i, output[i].ID, expectedOutput[i].ID)
		}
		if *output[i].Title != *expectedOutput[i].Title {
			t.Errorf("Unexpected title at index %d: got %v, want %v", i, *output[i].Title, *expectedOutput[i].Title)
		}
		if *output[i].Author != *expectedOutput[i].Author {
			t.Errorf("Unexpected author at index %d: got %v, want %v", i, *output[i].Author, *expectedOutput[i].Author)
		}
		if *output[i].BookCover != *expectedOutput[i].BookCover {
			t.Errorf("Unexpected book cover at index %d: got %v, want %v", i, *output[i].BookCover, *expectedOutput[i].BookCover)
		}
		if *output[i].PublishedAt != *expectedOutput[i].PublishedAt {
			t.Errorf("Unexpected published date at index %d: got %v, want %v", i, *output[i].PublishedAt, *expectedOutput[i].PublishedAt)
		}
	}

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetBook(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	// Create a new logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Create a new books service with the mock database connection and logger
	service := &bookssrvc{
		logger: logger,
		db:     db,
	}

	// Define the input and output for the test case
	id := 13
	title := "A Game of Thrones"
	author := "George R. R. Martin"
	bookCover := "https://www.example.com/a_game_of_thrones.jpg"
	publishedAt := "1996-08-01"

	expectedOutput := &books.Book{
		ID:          &id,
		Title:       &title,
		Author:      &author,
		BookCover:   &bookCover,
		PublishedAt: &publishedAt,
	}

	// Set up the mock database to expect the select query
	rows := sqlmock.NewRows([]string{"ID", "Title", "Author", "BookCover", "PublishedAt"}).
		AddRow(expectedOutput.ID, expectedOutput.Title, expectedOutput.Author, expectedOutput.BookCover, expectedOutput.PublishedAt)
	mock.ExpectQuery("^SELECT \\* FROM books WHERE Id = ?").WithArgs(id).WillReturnRows(rows)

	// Call the Get method on the books service
	output, err := service.GetBook(context.Background(), &books.GetBookPayload{ID: &id})

	// Check for errors and unexpected output
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if *output.ID != *expectedOutput.ID {
		t.Errorf("Unexpected ID: got %v, want %v", *output.ID, *expectedOutput.ID)
	}
	if *output.Title != *expectedOutput.Title {
		t.Errorf("Unexpected title: got %v, want %v", *output.Title, *expectedOutput.Title)
	}
	if *output.Author != *expectedOutput.Author {
		t.Errorf("Unexpected author: got %v, want %v", *output.Author, *expectedOutput.Author)
	}
	if *output.BookCover != *expectedOutput.BookCover {
		t.Errorf("Unexpected book cover: got %v, want %v", *output.BookCover, *expectedOutput.BookCover)
	}
	if *output.PublishedAt != *expectedOutput.PublishedAt {
		t.Errorf("Unexpected published date: got %v, want %v", *output.PublishedAt, *expectedOutput.PublishedAt)
	}
}
