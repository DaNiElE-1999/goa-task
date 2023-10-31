// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books client HTTP transport
//
// Command:
// $ goa gen books/design

package client

import (
	books "books/gen/books"
	"context"
	"mime/multipart"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the books service endpoint HTTP clients.
type Client struct {
	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// All Doer is the HTTP client used to make requests to the all endpoint.
	AllDoer goahttp.Doer

	// UpdateBook Doer is the HTTP client used to make requests to the updateBook
	// endpoint.
	UpdateBookDoer goahttp.Doer

	// GetBook Doer is the HTTP client used to make requests to the getBook
	// endpoint.
	GetBookDoer goahttp.Doer

	// DeleteBook Doer is the HTTP client used to make requests to the deleteBook
	// endpoint.
	DeleteBookDoer goahttp.Doer

	// UploadImage Doer is the HTTP client used to make requests to the uploadImage
	// endpoint.
	UploadImageDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// BooksUploadImageEncoderFunc is the type to encode multipart request for the
// "books" service "uploadImage" endpoint.
type BooksUploadImageEncoderFunc func(*multipart.Writer, *books.UploadImagePayload) error

// NewClient instantiates HTTP clients for all the books service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateDoer:          doer,
		AllDoer:             doer,
		UpdateBookDoer:      doer,
		GetBookDoer:         doer,
		DeleteBookDoer:      doer,
		UploadImageDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Create returns an endpoint that makes HTTP requests to the books service
// create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("books", "create", err)
		}
		return decodeResponse(resp)
	}
}

// All returns an endpoint that makes HTTP requests to the books service all
// server.
func (c *Client) All() goa.Endpoint {
	var (
		decodeResponse = DecodeAllResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAllRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AllDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("books", "all", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateBook returns an endpoint that makes HTTP requests to the books service
// updateBook server.
func (c *Client) UpdateBook() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateBookRequest(c.encoder)
		decodeResponse = DecodeUpdateBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("books", "updateBook", err)
		}
		return decodeResponse(resp)
	}
}

// GetBook returns an endpoint that makes HTTP requests to the books service
// getBook server.
func (c *Client) GetBook() goa.Endpoint {
	var (
		decodeResponse = DecodeGetBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("books", "getBook", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteBook returns an endpoint that makes HTTP requests to the books service
// deleteBook server.
func (c *Client) DeleteBook() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteBookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDeleteBookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteBookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("books", "deleteBook", err)
		}
		return decodeResponse(resp)
	}
}

// UploadImage returns an endpoint that makes HTTP requests to the books
// service uploadImage server.
func (c *Client) UploadImage(booksUploadImageEncoderFn BooksUploadImageEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeUploadImageRequest(NewBooksUploadImageEncoder(booksUploadImageEncoderFn))
		decodeResponse = DecodeUploadImageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUploadImageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UploadImageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("books", "uploadImage", err)
		}
		return decodeResponse(resp)
	}
}
