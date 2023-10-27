-- +goose Up
CREATE TABLE books (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Author VARCHAR(255) NOT NULL,
    BookCover VARCHAR(255),
    PublishedAt DATE
);
-- +goose Down
DROP TABLE books;