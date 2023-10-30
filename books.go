package booksapi

import (
	books "books/gen/books"
	context "context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Date format regular expression
var dateFormatRegexp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

// Custom error for invalid date format
var ErrInvalidDateFormat = errors.New("Invalid date format. Date should be in 'yyyy-mm-dd' format")

// Validate date format
func isValidDateFormat(date string) bool {
	return dateFormatRegexp.MatchString(date)
}

type bookssrvc struct {
	logger *log.Logger
	db     *sql.DB
	dir    string // Path to download and upload directory
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	// Create a database connection.
	godotenv.Load(".env")

	connectionString := os.Getenv("Connection_String")
	if connectionString == "" {
		log.Fatal("Connection string not found")
	}

	db, err := sql.Open("mysql", connectionString)
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

	// Validate the date format
	if p.PublishedAt != nil && !isValidDateFormat(*p.PublishedAt) {
		return nil, ErrInvalidDateFormat
	}

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

func (s *bookssrvc) UpdateBook(ctx context.Context, p *books.UpdateBookPayload) (res *books.Book, err error) {
	s.logger.Print("books.updateBook")

	// Validate the date format
	if p.Book.PublishedAt != nil && !isValidDateFormat(*p.Book.PublishedAt) {
		return nil, ErrInvalidDateFormat
	}

	// Initialize a list to store the SET clauses for the SQL query
	var setClauses []string

	// Initialize a list to store the arguments for the SQL query
	var args []interface{}

	// Check if the payload contains a title update
	if p.Book.Title != nil {
		setClauses = append(setClauses, "Title = ?")
		args = append(args, *p.Book.Title)
	}

	// Check if the payload contains an author update
	if p.Book.Author != nil {
		setClauses = append(setClauses, "Author = ?")
		args = append(args, *p.Book.Author)
	}

	// Check if the payload contains a bookCover update
	if p.Book.BookCover != nil {
		setClauses = append(setClauses, "BookCover = ?")
		args = append(args, *p.Book.BookCover)
	}

	// Check if the payload contains a publishedAt update
	if p.Book.PublishedAt != nil {
		setClauses = append(setClauses, "PublishedAt = ?")
		args = append(args, *p.Book.PublishedAt)
	}

	// Add the ID for the WHERE clause
	args = append(args, p.ID)

	// Create the SQL query with the SET clauses
	updateQuery := "UPDATE books SET " + strings.Join(setClauses, ", ") + " WHERE Id = ?"

	// Execute the query using the reusable function
	if err := s.executeQuery(updateQuery, args...); err != nil {
		return nil, err
	}

	// Fetch the updated book from the database
	row := s.db.QueryRowContext(ctx, "SELECT * FROM books WHERE Id = ?", p.ID)

	// Scan the row into a new Book struct
	var updatedBook books.Book
	if err := row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.BookCover, &updatedBook.PublishedAt); err != nil {
		return nil, err
	}

	// Return the updated book
	return &updatedBook, nil
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

// Upload implements file upload.
// func (s *bookssrvc) Upload(ctx context.Context, p *books.UploadPayload, req io.ReadCloser) error {
// 	defer req.Close() // Close the request body when done

// 	// Ensure the upload directory exists
// 	uploadDir := filepath.Join(s.dir, p.Dir)
// 	if err := os.MkdirAll(uploadDir, 0777); err != nil {
// 		return books.MakeInternalError(err)
// 	}

// 	s.logger.Printf("Upload directory: %s\n", uploadDir)

// 	// Create a multipart request reader
// 	_, params, err := mime.ParseMediaType(p.ContentType)
// 	if err != nil {
// 		return books.MakeInvalidMediaType(err)
// 	}
// 	mr := multipart.NewReader(req, params["boundary"])

// 	// Process each part of the multipart request
// 	for {
// 		part, err := mr.NextPart()
// 		if err == io.EOF {
// 			// No more parts, we're done!
// 			return nil
// 		}
// 		if err != nil {
// 			return books.MakeInvalidMultipartRequest(err)
// 		}

// 		// Create the uploaded file, potentially overwriting an existing file
// 		fpath := filepath.Join(uploadDir, part.FileName())
// 		f, err := os.Create(fpath)
// 		if err != nil {
// 			return books.MakeInternalError(err)
// 		}
// 		defer f.Close() // Close the file when done

// 		// Stream content to the disk
// 		n, copyErr := io.Copy(f, part)
// 		if copyErr != nil {
// 			return books.MakeInternalError(copyErr)
// 		}
// 		s.logger.Printf("Written %d bytes to %q", n, fpath)
// 	}
// }

// Upload implements upload.
func (s *bookssrvc) Upload(ctx context.Context, p *books.UploadPayload, req io.ReadCloser) error {
	// Don't forget to close the body reader!
	defer req.Close()

	s.logger.Print(s.dir)
	s.logger.Print(p.Dir)

	// Make sure upload directory exists
	uploadDir := filepath.Join(s.dir, p.Dir)
	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		return books.MakeInternalError(err)
	}

	// Createa multipart request reader
	_, params, err := mime.ParseMediaType(p.ContentType)
	if err != nil {
		return books.MakeInvalidMediaType(err)
	}
	mr := multipart.NewReader(req, params["boundary"])

	fmt.Printf("mr: %v\n", mr)

	//Go through each part and save the corresponding content to disk.
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			// We're done!
			return nil
		}
		if err != nil {
			return books.MakeInvalidMultipartRequest(err)
		}

		// Create uploaded file, potentially overridding existing file.
		fpath := filepath.Join(uploadDir, part.FileName())
		f, err := os.Create(fpath)
		if err != nil {
			return books.MakeInternalError(err)
		}
		defer f.Close()

		// Stream content to disk.
		n, err := io.Copy(f, part)
		if err != nil {
			return books.MakeInternalError(err)
		}
		s.logger.Printf("Written %d bytes to %q", n, fpath)
	}

}
