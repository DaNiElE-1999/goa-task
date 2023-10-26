// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books endpoints
//
// Command:
// $ goa gen books/design

package books

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "books" service endpoints.
type Endpoints struct {
	Create     goa.Endpoint
	All        goa.Endpoint
	UpdateBook goa.Endpoint
	GetBook    goa.Endpoint
	DeleteBook goa.Endpoint
}

// NewEndpoints wraps the methods of the "books" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create:     NewCreateEndpoint(s),
		All:        NewAllEndpoint(s),
		UpdateBook: NewUpdateBookEndpoint(s),
		GetBook:    NewGetBookEndpoint(s),
		DeleteBook: NewDeleteBookEndpoint(s),
	}
}

// Use applies the given middleware to all the "books" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.All = m(e.All)
	e.UpdateBook = m(e.UpdateBook)
	e.GetBook = m(e.GetBook)
	e.DeleteBook = m(e.DeleteBook)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "books".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*Book)
		return s.Create(ctx, p)
	}
}

// NewAllEndpoint returns an endpoint function that calls the method "all" of
// service "books".
func NewAllEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.All(ctx)
	}
}

// NewUpdateBookEndpoint returns an endpoint function that calls the method
// "updateBook" of service "books".
func NewUpdateBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateBookPayload)
		return s.UpdateBook(ctx, p)
	}
}

// NewGetBookEndpoint returns an endpoint function that calls the method
// "getBook" of service "books".
func NewGetBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetBookPayload)
		return s.GetBook(ctx, p)
	}
}

// NewDeleteBookEndpoint returns an endpoint function that calls the method
// "deleteBook" of service "books".
func NewDeleteBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteBookPayload)
		return nil, s.DeleteBook(ctx, p)
	}
}
