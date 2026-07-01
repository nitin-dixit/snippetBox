package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, pass string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return err
	}

	stmt := `insert into users (name,email, hashed_password,created)
	values ($1,$2,$3,current_timestamp)
	`
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "users_uc_email" {
				return ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, pass string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `select id,hashed_password from users where email=$1`

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(pass))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		}
		return 0, err
	}
	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
