-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS rate_books;

CREATE TABLE
    rate_books.users (
        id          UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        created_at  TIMESTAMPTZ DEFAULT now (),
        updated_at  TIMESTAMPTZ,
        name        VARCHAR NOT NULL,
        email       VARCHAR NOT NULL UNIQUE,
        pass        VARCHAR NOT NULL
    );
CREATE INDEX idx_users_name ON rate_books.users(name);
CREATE INDEX idx_users_email ON rate_books.users(email);

CREATE TABLE
    rate_books.authors (
        id          UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        created_at  TIMESTAMPTZ DEFAULT now (),
        updated_at  TIMESTAMPTZ,
        surname     VARCHAR NOT NULL,
        name        VARCHAR NOT NULL,
        country     VARCHAR,
        year_born   INTEGER,
        description VARCHAR
    );
CREATE INDEX idx_authors_name ON rate_books.authors(surname, name);
CREATE INDEX idx_authors_country ON rate_books.authors(country);

CREATE TABLE
    rate_books.genres (
        id          UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        created_at  TIMESTAMPTZ DEFAULT now (),
        updated_at  TIMESTAMPTZ,
        name        VARCHAR NOT NULL UNIQUE
    );
CREATE INDEX idx_genres_name ON rate_books.genres(name);

CREATE TABLE
    rate_books.books (
        id          UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        created_at  TIMESTAMPTZ DEFAULT now (),
        updated_at  TIMESTAMPTZ,
        title       VARCHAR NOT NULL,
        author_id   UUID REFERENCES rate_books.authors(id) ON DELETE CASCADE,
        owner_id    UUID REFERENCES rate_books.users(id),
        year_public INTEGER,
        year_read   INTEGER NOT NULL,
        grade       INTEGER NOT NULL,
        description VARCHAR
    );
CREATE INDEX idx_books_title ON rate_books.books(title);
CREATE INDEX idx_books_author_id ON rate_books.books(author_id);

CREATE TABLE rate_books.book_genres (
    book_id     UUID REFERENCES rate_books.books(id) ON DELETE CASCADE,
    genre_id    UUID REFERENCES rate_books.genres(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, genre_id)
);
CREATE INDEX idx_book_genres_book ON rate_books.book_genres(book_id);
CREATE INDEX idx_book_genres_genre ON rate_books.book_genres(genre_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rate_books.book_genres;
DROP TABLE IF EXISTS rate_books.books;
DROP TABLE IF EXISTS rate_books.genres;
DROP TABLE IF EXISTS rate_books.authors;
DROP TABLE IF EXISTS rate_books.users;
DROP SCHEMA IF EXISTS rate_books;
-- +goose StatementEnd
