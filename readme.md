# Setup the project

    After cloning the project:

1. Install latest golang version
2. Install the latest goa framework

   go install goa.design/goa/v3/cmd/goa@v3
3. Install protoc using the documentation, according to your OS distribution
4. Install sqlc using the documentation, according to your OS distribution
5. Install goose using the documentation, according to your OS distribution
6. Install all the rest of packages:

   go get ./...

# Setup the database

1. Open a shell session on the schema directory (.\\sql\schema\)
2. Type: "goose mysql 'Connection_String' up"

This will run all the migrations.

# API

- ###### Create

Send a JSON POST request to  "/books" endpoint with the following key-values:

'{
      "author": "Sunt ut sint accusamus.",
      "bookCover": "Omnis molestiae sed.",
      "publishedAt": "yyyy-mm-dd",
      "title": "Ipsam sed."
   }'

- ###### All

Send a GET request to "/books" endpoint, and you will be returned a json list of books:

- ###### Update Book

Send a JSON PUT request to  "/books/{id}" endpoint with the following key-values:

{
  "book": {
      "author": "Sunt ut sint accusamus.",
      "bookCover": "Omnis molestiae sed.",
      "publishedAt": "yyyy-mm-dd",
      "title": "Ipsam sed."
  }
}

Please note that the update request must contain at least one value

- ###### Get Book

Send a GET request to "/books/{id}" endpoint, and you will be returned the book with that ID:

- ###### Delete Book

Send a DELETE request to "/books/{id}" endpoint, to delete the book with that ID

- ###### Upload

ToBeImplemented

# Note

Normally instead of the book cover path, I would have used a foreign key pointing to the bookCovers table. Since I do not have a frontend, I am using this method, even though not ideal.
