package booksapi

import (
	"books/gen/books"
	"mime/multipart"
)

// BooksUploadImageDecoderFunc implements the multipart decoder for service
// "books" endpoint "uploadImage". The decoder must populate the argument p
// after encoding.
func BooksUploadImageDecoderFunc(mr *multipart.Reader, p **books.UploadImagePayload) error {
	// Add multipart request decoder logic here
	return nil
}

// BooksUploadImageEncoderFunc implements the multipart encoder for service
// "books" endpoint "uploadImage".
func BooksUploadImageEncoderFunc(mw *multipart.Writer, p *books.UploadImagePayload) error {
	// Add multipart request encoder logic here
	return nil
}
