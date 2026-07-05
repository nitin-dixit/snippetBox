package models

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pressly/goose/v3"
	"golang.org/x/crypto/bcrypt"
)

func newTestDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal("Error loading .env file", err)
	}
	pool, err := pgxpool.New(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}

	db := stdlib.OpenDBFromPool(pool)

	// run migrations
	err = goose.SetDialect("postgres")
	if err != nil {
		t.Fatal(err)
	}

	err = goose.Up(db, "../../migrations/")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`TRUNCATE users restart identity CASCADE`)
	if err != nil {
		t.Fatalf("truncating tables %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("pa55word"),
		12,
	)
	if err != nil {
		t.Fatal(err)
	}

	stmt := `INSERT INTO users (name, email, hashed_password,created)
		VALUES ($1, $2, $3, $4)`

	_, err = pool.Exec(context.Background(), stmt, "Alice Jones",
		"alice@example.com",
		string(hashedPassword),
		time.Now())
	if err != nil {
		t.Fatal(err)
	}
	return db
}
