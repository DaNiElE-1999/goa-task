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
		Error("bad_request", func() {
			Description("The request is invalid")
		})
		HTTP(func() {
			POST("/books")
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("all", func() {
		Result(ArrayOf(Book))
		Error("not_found", func() {
			Description("Books not found")
		})
		HTTP(func() {
			GET("/books")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("updateBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
			Field(2, "book", Book)
		})
		Result(Book)
		Error("bad_request", func() {
			Description("The request is invalid")
		})
		HTTP(func() {
			PUT("/books/{id}")
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("getBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		Result(Book)
		Error("not_found", func() {
			Description("Book not found")
		})
		HTTP(func() {
			GET("/books/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("deleteBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		Error("not_found", func() {
			Description("Book not found")
		})
		HTTP(func() {
			DELETE("/books/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})

	Method("uploadImage", func() {
		Description("Upload an image")
		Payload(func() {
			Attribute("image", Bytes, "Binary data of the image")
			Attribute("content_type", String, "Content-Type header, must define value for multipart boundary.", func() {
				Default("multipart/form-data; boundary=goa")
				Pattern("multipart/[^;]+; boundary=.+")
			})
		})
		Error("bad_request", func() {
			Description("The request is invalid")
		})
		HTTP(func() {
			POST("/uploadBookCover")
			MultipartRequest()
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})

	Files("/openapi3.json", "./gen/http/openapi3.json")
})
