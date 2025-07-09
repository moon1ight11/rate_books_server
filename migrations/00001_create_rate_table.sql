-- +goose Up
-- +goose StatementBegin

CREATE TABLE authors (
  id SERIAL PRIMARY KEY,
  author_name VARCHAR,
  year_b INTEGER,
  country VARCHAR
);

CREATE TABLE rate_books (
  id SERIAL PRIMARY KEY,
  title VARCHAR UNIQUE,
  author_id INTEGER references authors(id),
  year_public INTEGER,
  year_read INTEGER,
  rate INTEGER,
  time_stamp timestamp DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rate_books;
DROP TABLE authors;
-- +goose StatementEnd
