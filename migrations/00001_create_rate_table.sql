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

CREATE TABLE covers (
  id SERIAL PRIMARY KEY,
  original_name VARCHAR,
  created_at timestamp DEFAULT now()
);

CREATE TABLE rate_books (
  id SERIAL PRIMARY KEY,
  title VARCHAR UNIQUE,
  author_id INTEGER references authors(id),
  year_public INTEGER,
  year_read INTEGER,
  rate INTEGER,
  cover_id INTEGER references covers(id),
  book_owner INTEGER references users(id),
  time_stamp timestamp DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE authors;
DROP TABLE rate_books;
-- +goose StatementEnd
