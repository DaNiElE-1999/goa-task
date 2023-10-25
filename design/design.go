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
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("books", func() {
	Description("The books service performs CRUD operations.")
	//
	HTTP(func() {
		Path("/library")
		Error("not_found", StatusNotFound, "Resource not found")
	})

	Method("listBooks", func() {
		Result(CollectionOf(Book))
		HTTP(func() {
			GET("/books")
		})
		GRPC(func() {
		})
	})

	Method("getBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		Result(Book)
		HTTP(func() {
			GET("/books/{id}")
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
		})
	})

	Method("createBook", func() {
		Payload(Book)
		Result(Book)
		HTTP(func() {
			POST("/books")
		})
		GRPC(func() {
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
		GRPC(func() {
		})
	})

	Method("deleteBook", func() {
		Payload(func() {
			Field(1, "id", Int, "Book ID")
		})
		HTTP(func() {
			DELETE("/books/{id}")
		})
		GRPC(func() {
		})
	})
	//
	Files("/openapi.json", "./gen/http/openapi.json")
})
