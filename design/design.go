package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("books", func() {
	Title("Books Service")
	Description("CRUD Application for books")
	Server("books", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var Book = Type("Book", func() {
	Attribute("ID", Int, "Unique ID of the book")
	Attribute("title", String, "Title of the book")
	Attribute("author", String, "Author of the book")
	Attribute("bookCover", String, "Cover of the book")
	Attribute("publishedAt", String, "Date the book has been published")
	Required("title", "author", "bookCover", "publishedAt")
})

var BookCover = Type("BookCover", func() {
	Attribute("bookTitle", String, "Title of the book")
	Attribute("imagePath", String, "Path to the image in storage")
})

var _ = Service("books", func() {
	Description("API for Books")

	Method("create", func() {
		Payload(Book)
		Result(Book)
		HTTP(func() {
			POST("/books")
		})
	})

	Method("all", func() {
		Result(ArrayOf(Book))
		HTTP(func() {
			GET("/books")
		})
	})

	Method("updateBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
			Field(2, "book", Book)
		})
		Result(Book)
		HTTP(func() {
			PUT("/books/{id}")
		})
	})

	Method("getBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		Result(Book)
		HTTP(func() {
			GET("/books/{id}")
		})
	})

	Method("deleteBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		HTTP(func() {
			DELETE("/books/{id}")
		})
	})

	Method("upload", func() {

		// The payload defines the request headers and parameters. It cannot
		// define body attributes as the endpoint makes use of
		// SkipRequestBodyEncodeDecode.
		Payload(func() {
			Attribute("content_type", String, "Content-Type header, must define value for multipart boundary.", func() {
				Default("multipart/form-data; boundary=goa")
				Pattern("multipart/[^;]+; boundary=.+")
				Example("multipart/form-data; boundary=goa")
			})
			Attribute("dir", String, "Dir is the relative path to the file directory where the uploaded content is saved.", func() {
				Default("upload")
				Example("upload")
			})
		})

		Error("invalid_media_type", ErrorResult, "Error returned when the Content-Type header does not define a multipart request.")
		Error("invalid_multipart_request", ErrorResult, "Error returned when the request body is not a valid multipart content.")
		Error("internal_error", ErrorResult, "Fault while processing upload.")

		HTTP(func() {
			POST("/upload/{*dir}")
			Header("content_type:Content-Type")

			// Bypass request body decoder code generation to alleviate need for
			// loading the entire request body in memory. The service gets
			// direct access to the HTTP request body reader.
			SkipRequestBodyEncodeDecode()

			// Define error HTTP statuses.
			Response("invalid_media_type", StatusBadRequest)
			Response("invalid_multipart_request", StatusBadRequest)
			Response("internal_error", StatusInternalServerError)
		})
	})

	Files("/openapi3.json", "./gen/http/openapi3.json")
}) //error handling to be implemented
