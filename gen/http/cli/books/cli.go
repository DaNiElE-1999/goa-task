// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books HTTP client CLI support package
//
// Command:
// $ goa gen books/design

package cli

import (
	booksc "books/gen/http/books/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `books (create|all|update-book|get-book|delete-book)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` books create --body '{
      "ID": 4698454743706616641,
      "author": "Ipsam sed.",
      "publishedAt": "Sunt ut sint accusamus.",
      "title": "Et odio est."
   }' --book-cover "multipart/form-data; boundary=goa"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		booksFlags = flag.NewFlagSet("books", flag.ContinueOnError)

		booksCreateFlags         = flag.NewFlagSet("create", flag.ExitOnError)
		booksCreateBodyFlag      = booksCreateFlags.String("body", "REQUIRED", "")
		booksCreateBookCoverFlag = booksCreateFlags.String("book-cover", "multipart/form-data; boundary=goa", "")

		booksAllFlags = flag.NewFlagSet("all", flag.ExitOnError)

		booksUpdateBookFlags    = flag.NewFlagSet("update-book", flag.ExitOnError)
		booksUpdateBookBodyFlag = booksUpdateBookFlags.String("body", "REQUIRED", "")
		booksUpdateBookIDFlag   = booksUpdateBookFlags.String("id", "REQUIRED", "Book ID")

		booksGetBookFlags  = flag.NewFlagSet("get-book", flag.ExitOnError)
		booksGetBookIDFlag = booksGetBookFlags.String("id", "REQUIRED", "Book ID")

		booksDeleteBookFlags  = flag.NewFlagSet("delete-book", flag.ExitOnError)
		booksDeleteBookIDFlag = booksDeleteBookFlags.String("id", "REQUIRED", "Book ID")
	)
	booksFlags.Usage = booksUsage
	booksCreateFlags.Usage = booksCreateUsage
	booksAllFlags.Usage = booksAllUsage
	booksUpdateBookFlags.Usage = booksUpdateBookUsage
	booksGetBookFlags.Usage = booksGetBookUsage
	booksDeleteBookFlags.Usage = booksDeleteBookUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "books":
			svcf = booksFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "books":
			switch epn {
			case "create":
				epf = booksCreateFlags

			case "all":
				epf = booksAllFlags

			case "update-book":
				epf = booksUpdateBookFlags

			case "get-book":
				epf = booksGetBookFlags

			case "delete-book":
				epf = booksDeleteBookFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "books":
			c := booksc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = booksc.BuildCreatePayload(*booksCreateBodyFlag, *booksCreateBookCoverFlag)
			case "all":
				endpoint = c.All()
				data = nil
			case "update-book":
				endpoint = c.UpdateBook()
				data, err = booksc.BuildUpdateBookPayload(*booksUpdateBookBodyFlag, *booksUpdateBookIDFlag)
			case "get-book":
				endpoint = c.GetBook()
				data, err = booksc.BuildGetBookPayload(*booksGetBookIDFlag)
			case "delete-book":
				endpoint = c.DeleteBook()
				data, err = booksc.BuildDeleteBookPayload(*booksDeleteBookIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// booksUsage displays the usage of the books command and its subcommands.
func booksUsage() {
	fmt.Fprintf(os.Stderr, `API for Books
Usage:
    %[1]s [globalflags] books COMMAND [flags]

COMMAND:
    create: Create implements create.
    all: All implements all.
    update-book: UpdateBook implements updateBook.
    get-book: GetBook implements getBook.
    delete-book: DeleteBook implements deleteBook.

Additional help:
    %[1]s books COMMAND --help
`, os.Args[0])
}
func booksCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books create -body JSON -book-cover STRING

Create implements create.
    -body JSON: 
    -book-cover STRING: 

Example:
    %[1]s books create --body '{
      "ID": 4698454743706616641,
      "author": "Ipsam sed.",
      "publishedAt": "Sunt ut sint accusamus.",
      "title": "Et odio est."
   }' --book-cover "multipart/form-data; boundary=goa"
`, os.Args[0])
}

func booksAllUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books all

All implements all.

Example:
    %[1]s books all
`, os.Args[0])
}

func booksUpdateBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books update-book -body JSON -id INT

UpdateBook implements updateBook.
    -body JSON: 
    -id INT: Book ID

Example:
    %[1]s books update-book --body '{
      "book": {
         "ID": 3012804354026031466,
         "author": "Aut eligendi repudiandae repellat.",
         "bookCover": "multipart/form-data; boundary=goa",
         "publishedAt": "Qui ut et quo est omnis dolor.",
         "title": "Mollitia assumenda explicabo impedit nesciunt."
      }
   }' --id 4485231234075589164
`, os.Args[0])
}

func booksGetBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books get-book -id INT

GetBook implements getBook.
    -id INT: Book ID

Example:
    %[1]s books get-book --id 9047210969149619277
`, os.Args[0])
}

func booksDeleteBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books delete-book -id INT

DeleteBook implements deleteBook.
    -id INT: Book ID

Example:
    %[1]s books delete-book --id 7164475420641560117
`, os.Args[0])
}
