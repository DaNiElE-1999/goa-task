// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books HTTP server encoders and decoders
//
// Command:
// $ goa gen books/design

package server

import (
	books "books/gen/books"
	"context"
	"io"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the books
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*books.Book)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the books create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateBook(&body)

		return payload, nil
	}
}

// EncodeAllResponse returns an encoder for responses returned by the books all
// endpoint.
func EncodeAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*books.Book)
		enc := encoder(ctx, w)
		body := NewAllResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeUpdateBookResponse returns an encoder for responses returned by the
// books updateBook endpoint.
func EncodeUpdateBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*books.Book)
		enc := encoder(ctx, w)
		body := NewUpdateBookResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateBookRequest returns a decoder for requests sent to the books
// updateBook endpoint.
func DecodeUpdateBookRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateBookRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateBookRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id int

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateBookPayload(&body, id)

		return payload, nil
	}
}

// EncodeGetBookResponse returns an encoder for responses returned by the books
// getBook endpoint.
func EncodeGetBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*books.Book)
		enc := encoder(ctx, w)
		body := NewGetBookResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetBookRequest returns a decoder for requests sent to the books
// getBook endpoint.
func DecodeGetBookRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetBookPayload(id)

		return payload, nil
	}
}

// EncodeDeleteBookResponse returns an encoder for responses returned by the
// books deleteBook endpoint.
func EncodeDeleteBookResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteBookRequest returns a decoder for requests sent to the books
// deleteBook endpoint.
func DecodeDeleteBookRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteBookPayload(id)

		return payload, nil
	}
}

// marshalBooksBookToBookResponse builds a value of type *BookResponse from a
// value of type *books.Book.
func marshalBooksBookToBookResponse(v *books.Book) *BookResponse {
	res := &BookResponse{
		ID:          v.ID,
		Title:       v.Title,
		Author:      v.Author,
		BookCover:   v.BookCover,
		PublishedAt: v.PublishedAt,
	}

	return res
}

// unmarshalBookRequestBodyToBooksBook builds a value of type *books.Book from
// a value of type *BookRequestBody.
func unmarshalBookRequestBodyToBooksBook(v *BookRequestBody) *books.Book {
	if v == nil {
		return nil
	}
	res := &books.Book{
		ID:          v.ID,
		Title:       *v.Title,
		Author:      *v.Author,
		BookCover:   *v.BookCover,
		PublishedAt: *v.PublishedAt,
	}

	return res
}
