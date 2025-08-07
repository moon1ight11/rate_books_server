package database

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

// объявление УРЛ для БД и директории миграций
const (
	url           = "postgres://fedor:fedor_pass@localhost:15432/rate"
	migrationsDir = "./migrations"
)

var DB *sql.DB

// соединение с БД
func Connection() {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	DB = stdlib.OpenDB(*config)

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

// обновление миграций
func UppingMigrations() {
	goose.SetBaseFS(nil)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	err := goose.Up(DB, migrationsDir)
	if err != nil {
		log.Fatalf("Failed to up migrations: %v", err)
	}
}
