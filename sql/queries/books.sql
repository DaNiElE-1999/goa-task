-- name: CreateBook :one
INSERT INTO books (Id, Title, Author, BookCover, PublishedAt)
VALUES (?, ?, ?, ?, ?);
RETURNING *;