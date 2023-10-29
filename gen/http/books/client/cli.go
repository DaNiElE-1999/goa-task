// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books HTTP client CLI support package
//
// Command:
// $ goa gen books/design

package client

import (
	books "books/gen/books"
	"encoding/json"
	"fmt"
	"strconv"

	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the books create endpoint from CLI
// flags.
func BuildCreatePayload(booksCreateBody string, booksCreateBookCover string) (*books.Book, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(booksCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"ID\": 4698454743706616641,\n      \"author\": \"Ipsam sed.\",\n      \"publishedAt\": \"Sunt ut sint accusamus.\",\n      \"title\": \"Et odio est.\"\n   }'")
		}
	}
	var bookCover string
	{
		bookCover = booksCreateBookCover
		err = goa.MergeErrors(err, goa.ValidatePattern("bookCover", bookCover, "multipart/[^;]+; boundary=.+"))
		if err != nil {
			return nil, err
		}
	}
	v := &books.Book{
		ID:          body.ID,
		Title:       body.Title,
		Author:      body.Author,
		PublishedAt: body.PublishedAt,
	}
	v.BookCover = bookCover

	return v, nil
}

// BuildUpdateBookPayload builds the payload for the books updateBook endpoint
// from CLI flags.
func BuildUpdateBookPayload(booksUpdateBookBody string, booksUpdateBookID string) (*books.UpdateBookPayload, error) {
	var err error
	var body UpdateBookRequestBody
	{
		err = json.Unmarshal([]byte(booksUpdateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"book\": {\n         \"ID\": 3012804354026031466,\n         \"author\": \"Aut eligendi repudiandae repellat.\",\n         \"bookCover\": \"multipart/form-data; boundary=goa\",\n         \"publishedAt\": \"Qui ut et quo est omnis dolor.\",\n         \"title\": \"Mollitia assumenda explicabo impedit nesciunt.\"\n      }\n   }'")
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(booksUpdateBookID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &books.UpdateBookPayload{}
	if body.Book != nil {
		v.Book = marshalBookRequestBodyToBooksBook(body.Book)
	}
	v.ID = &id

	return v, nil
}

// BuildGetBookPayload builds the payload for the books getBook endpoint from
// CLI flags.
func BuildGetBookPayload(booksGetBookID string) (*books.GetBookPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(booksGetBookID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &books.GetBookPayload{}
	v.ID = &id

	return v, nil
}

// BuildDeleteBookPayload builds the payload for the books deleteBook endpoint
// from CLI flags.
func BuildDeleteBookPayload(booksDeleteBookID string) (*books.DeleteBookPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(booksDeleteBookID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &books.DeleteBookPayload{}
	v.ID = &id

	return v, nil
}
