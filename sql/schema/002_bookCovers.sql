-- +goose Up
CREATE TABLE bookCovers (
    Id INT AUTO_INCREMENT PRIMARY KEY,
    Path VARCHAR(255) NOT NULL,
    BookId INT,
    FOREIGN KEY (BookId) REFERENCES books(Id)
);
-- +goose Down
DROP TABLE bookCovers;