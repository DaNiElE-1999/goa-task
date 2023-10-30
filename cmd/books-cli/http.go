package main

import (
	books "books"
	cli "books/gen/http/cli/books"
	"net/http"
	"time"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

func doHTTP(scheme, host string, timeout int, debug bool) (goa.Endpoint, any, error) {
	var (
		doer goahttp.Doer
	)
	{
		doer = &http.Client{Timeout: time.Duration(timeout) * time.Second}
		if debug {
			doer = goahttp.NewDebugDoer(doer)
		}
	}

	return cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		debug,
		books.BooksUploadImageEncoderFunc,
	)
}

func httpUsageCommands() string {
	return cli.UsageCommands()
}

func httpUsageExamples() string {
	return cli.UsageExamples()
}
