// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books HTTP client encoders and decoders
//
// Command:
// $ goa gen books/design

package client

import (
	books "books/gen/books"
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "books" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateBooksPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the books create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*books.Book)
		if !ok {
			return goahttp.ErrInvalidType("books", "create", "*books.Book", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("books", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the books
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "create", err)
			}
			res := NewCreateBookOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildAllRequest instantiates a HTTP request object with method and path set
// to call the "books" service "all" endpoint
func (c *Client) BuildAllRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AllBooksPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "all", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeAllResponse returns a decoder for responses returned by the books all
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AllResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "all", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateBookResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "all", err)
			}
			res := NewAllBookOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "all", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateBookRequest instantiates a HTTP request object with method and
// path set to call the "books" service "updateBook" endpoint
func (c *Client) BuildUpdateBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*books.UpdateBookPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("books", "updateBook", "*books.UpdateBookPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateBookBooksPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "updateBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateBookRequest returns an encoder for requests sent to the books
// updateBook server.
func EncodeUpdateBookRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*books.UpdateBookPayload)
		if !ok {
			return goahttp.ErrInvalidType("books", "updateBook", "*books.UpdateBookPayload", v)
		}
		body := NewUpdateBookRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("books", "updateBook", err)
		}
		return nil
	}
}

// DecodeUpdateBookResponse returns a decoder for responses returned by the
// books updateBook endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUpdateBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateBookResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "updateBook", err)
			}
			err = ValidateUpdateBookResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "updateBook", err)
			}
			res := NewUpdateBookBookOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "updateBook", resp.StatusCode, string(body))
		}
	}
}

// BuildGetBookRequest instantiates a HTTP request object with method and path
// set to call the "books" service "getBook" endpoint
func (c *Client) BuildGetBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*books.GetBookPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("books", "getBook", "*books.GetBookPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetBookBooksPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "getBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetBookResponse returns a decoder for responses returned by the books
// getBook endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeGetBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetBookResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("books", "getBook", err)
			}
			err = ValidateGetBookResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("books", "getBook", err)
			}
			res := NewGetBookBookOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "getBook", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteBookRequest instantiates a HTTP request object with method and
// path set to call the "books" service "deleteBook" endpoint
func (c *Client) BuildDeleteBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*books.DeleteBookPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("books", "deleteBook", "*books.DeleteBookPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteBookBooksPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "deleteBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteBookResponse returns a decoder for responses returned by the
// books deleteBook endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeDeleteBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "deleteBook", resp.StatusCode, string(body))
		}
	}
}

// BuildUploadImageRequest instantiates a HTTP request object with method and
// path set to call the "books" service "uploadImage" endpoint
func (c *Client) BuildUploadImageRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UploadImageBooksPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("books", "uploadImage", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUploadImageRequest returns an encoder for requests sent to the books
// uploadImage server.
func EncodeUploadImageRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*books.UploadImagePayload)
		if !ok {
			return goahttp.ErrInvalidType("books", "uploadImage", "*books.UploadImagePayload", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("books", "uploadImage", err)
		}
		return nil
	}
}

// NewBooksUploadImageEncoder returns an encoder to encode the multipart
// request for the "books" service "uploadImage" endpoint.
func NewBooksUploadImageEncoder(encoderFn BooksUploadImageEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v any) error {
			p := v.(*books.UploadImagePayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = io.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeUploadImageResponse returns a decoder for responses returned by the
// books uploadImage endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUploadImageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("books", "uploadImage", resp.StatusCode, string(body))
		}
	}
}

// unmarshalBookResponseToBooksBook builds a value of type *books.Book from a
// value of type *BookResponse.
func unmarshalBookResponseToBooksBook(v *BookResponse) *books.Book {
	res := &books.Book{
		ID:          v.ID,
		Title:       *v.Title,
		Author:      *v.Author,
		BookCover:   *v.BookCover,
		PublishedAt: *v.PublishedAt,
	}

	return res
}

// marshalBooksBookToBookRequestBody builds a value of type *BookRequestBody
// from a value of type *books.Book.
func marshalBooksBookToBookRequestBody(v *books.Book) *BookRequestBody {
	if v == nil {
		return nil
	}
	res := &BookRequestBody{
		ID:          v.ID,
		Title:       v.Title,
		Author:      v.Author,
		BookCover:   v.BookCover,
		PublishedAt: v.PublishedAt,
	}

	return res
}

// marshalBookRequestBodyToBooksBook builds a value of type *books.Book from a
// value of type *BookRequestBody.
func marshalBookRequestBodyToBooksBook(v *BookRequestBody) *books.Book {
	if v == nil {
		return nil
	}
	res := &books.Book{
		ID:          v.ID,
		Title:       v.Title,
		Author:      v.Author,
		BookCover:   v.BookCover,
		PublishedAt: v.PublishedAt,
	}

	return res
}
