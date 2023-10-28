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
	Attribute("title", String, "Title of the book", func() {
		Required()
	})
	Attribute("author", String, "Author of the book", func() {
		Required()
	})
	Attribute("bookCover", String, "Cover of the book", func() {
		Required()
	})
	Attribute("publishedAt", String, "Date the book has been published", func() {
		Required()
	})
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

	Files("/openapi3.json", "./gen/http/openapi3.json")
}) //error handling to be implemented
