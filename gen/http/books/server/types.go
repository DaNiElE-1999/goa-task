// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books HTTP server types
//
// Command:
// $ goa gen books/design

package server

import (
	books "books/gen/books"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "books" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Unique ID of the book
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Cover of the book
	BookCover *string `form:"bookCover,omitempty" json:"bookCover,omitempty" xml:"bookCover,omitempty"`
	// Date the book has been published
	PublishedAt *string `form:"publishedAt,omitempty" json:"publishedAt,omitempty" xml:"publishedAt,omitempty"`
}

// UpdateBookRequestBody is the type of the "books" service "updateBook"
// endpoint HTTP request body.
type UpdateBookRequestBody struct {
	Book *BookRequestBody `form:"book,omitempty" json:"book,omitempty" xml:"book,omitempty"`
}

// UploadImageRequestBody is the type of the "books" service "uploadImage"
// endpoint HTTP request body.
type UploadImageRequestBody struct {
	// Binary data of the image
	Image []byte `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Content-Type header, must define value for multipart boundary.
	ContentType *string `form:"content_type,omitempty" json:"content_type,omitempty" xml:"content_type,omitempty"`
}

// CreateResponseBody is the type of the "books" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// Unique ID of the book
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Cover of the book
	BookCover string `form:"bookCover" json:"bookCover" xml:"bookCover"`
	// Date the book has been published
	PublishedAt string `form:"publishedAt" json:"publishedAt" xml:"publishedAt"`
}

// AllResponseBody is the type of the "books" service "all" endpoint HTTP
// response body.
type AllResponseBody []*BookResponse

// UpdateBookResponseBody is the type of the "books" service "updateBook"
// endpoint HTTP response body.
type UpdateBookResponseBody struct {
	// Unique ID of the book
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Cover of the book
	BookCover string `form:"bookCover" json:"bookCover" xml:"bookCover"`
	// Date the book has been published
	PublishedAt string `form:"publishedAt" json:"publishedAt" xml:"publishedAt"`
}

// GetBookResponseBody is the type of the "books" service "getBook" endpoint
// HTTP response body.
type GetBookResponseBody struct {
	// Unique ID of the book
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Cover of the book
	BookCover string `form:"bookCover" json:"bookCover" xml:"bookCover"`
	// Date the book has been published
	PublishedAt string `form:"publishedAt" json:"publishedAt" xml:"publishedAt"`
}

// BookResponse is used to define fields on response body types.
type BookResponse struct {
	// Unique ID of the book
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Title of the book
	Title string `form:"title" json:"title" xml:"title"`
	// Author of the book
	Author string `form:"author" json:"author" xml:"author"`
	// Cover of the book
	BookCover string `form:"bookCover" json:"bookCover" xml:"bookCover"`
	// Date the book has been published
	PublishedAt string `form:"publishedAt" json:"publishedAt" xml:"publishedAt"`
}

// BookRequestBody is used to define fields on request body types.
type BookRequestBody struct {
	// Unique ID of the book
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Title of the book
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Author of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Cover of the book
	BookCover *string `form:"bookCover,omitempty" json:"bookCover,omitempty" xml:"bookCover,omitempty"`
	// Date the book has been published
	PublishedAt *string `form:"publishedAt,omitempty" json:"publishedAt,omitempty" xml:"publishedAt,omitempty"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "books" service.
func NewCreateResponseBody(res *books.Book) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:          res.ID,
		Title:       res.Title,
		Author:      res.Author,
		BookCover:   res.BookCover,
		PublishedAt: res.PublishedAt,
	}
	return body
}

// NewAllResponseBody builds the HTTP response body from the result of the
// "all" endpoint of the "books" service.
func NewAllResponseBody(res []*books.Book) AllResponseBody {
	body := make([]*BookResponse, len(res))
	for i, val := range res {
		body[i] = marshalBooksBookToBookResponse(val)
	}
	return body
}

// NewUpdateBookResponseBody builds the HTTP response body from the result of
// the "updateBook" endpoint of the "books" service.
func NewUpdateBookResponseBody(res *books.Book) *UpdateBookResponseBody {
	body := &UpdateBookResponseBody{
		ID:          res.ID,
		Title:       res.Title,
		Author:      res.Author,
		BookCover:   res.BookCover,
		PublishedAt: res.PublishedAt,
	}
	return body
}

// NewGetBookResponseBody builds the HTTP response body from the result of the
// "getBook" endpoint of the "books" service.
func NewGetBookResponseBody(res *books.Book) *GetBookResponseBody {
	body := &GetBookResponseBody{
		ID:          res.ID,
		Title:       res.Title,
		Author:      res.Author,
		BookCover:   res.BookCover,
		PublishedAt: res.PublishedAt,
	}
	return body
}

// NewCreateBook builds a books service create endpoint payload.
func NewCreateBook(body *CreateRequestBody) *books.Book {
	v := &books.Book{
		ID:          body.ID,
		Title:       *body.Title,
		Author:      *body.Author,
		BookCover:   *body.BookCover,
		PublishedAt: *body.PublishedAt,
	}

	return v
}

// NewUpdateBookPayload builds a books service updateBook endpoint payload.
func NewUpdateBookPayload(body *UpdateBookRequestBody, id int) *books.UpdateBookPayload {
	v := &books.UpdateBookPayload{}
	if body.Book != nil {
		v.Book = unmarshalBookRequestBodyToBooksBook(body.Book)
	}
	v.ID = &id

	return v
}

// NewGetBookPayload builds a books service getBook endpoint payload.
func NewGetBookPayload(id int) *books.GetBookPayload {
	v := &books.GetBookPayload{}
	v.ID = &id

	return v
}

// NewDeleteBookPayload builds a books service deleteBook endpoint payload.
func NewDeleteBookPayload(id int) *books.DeleteBookPayload {
	v := &books.DeleteBookPayload{}
	v.ID = &id

	return v
}

// NewUploadImagePayload builds a books service uploadImage endpoint payload.
func NewUploadImagePayload(body *UploadImageRequestBody) *books.UploadImagePayload {
	v := &books.UploadImagePayload{
		Image: body.Image,
	}
	if body.ContentType != nil {
		v.ContentType = *body.ContentType
	}
	if body.ContentType == nil {
		v.ContentType = "multipart/form-data; boundary=goa"
	}

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bookCover", "body"))
	}
	if body.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("publishedAt", "body"))
	}
	return
}

// ValidateUpdateBookRequestBody runs the validations defined on
// UpdateBookRequestBody
func ValidateUpdateBookRequestBody(body *UpdateBookRequestBody) (err error) {
	if body.Book != nil {
		if err2 := ValidateBookRequestBody(body.Book); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUploadImageRequestBody runs the validations defined on
// UploadImageRequestBody
func ValidateUploadImageRequestBody(body *UploadImageRequestBody) (err error) {
	if body.ContentType != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.content_type", *body.ContentType, "multipart/[^;]+; boundary=.+"))
	}
	return
}

// ValidateBookRequestBody runs the validations defined on BookRequestBody
func ValidateBookRequestBody(body *BookRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bookCover", "body"))
	}
	if body.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("publishedAt", "body"))
	}
	return
}
