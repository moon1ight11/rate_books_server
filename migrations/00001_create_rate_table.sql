-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  user_name VARCHAR,
  user_password VARCHAR,
  time_stamp timestamp DEFAULT now()
);

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
  time_stamp timestamp DEFAULT now(),
  book_owner INTEGER references users(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE authors;
DROP TABLE rate_books;
-- +goose StatementEnd
