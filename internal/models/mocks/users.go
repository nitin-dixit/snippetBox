package mocks

import (
	"time"

	"github.com/nitin-dixit/snippetBox/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, pass string) error {
	switch email {
	case "demo@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, pass string) (int, error) {
	if email == "alice@example.com" && pass == "pa$$word" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (models.User, error) {
	if id == 1 {
		u := models.User{
			ID:      1,
			Name:    "Alice",
			Email:   "alice@example.com",
			Created: time.Now(),
		}
		return u, nil
	}

	return models.User{}, models.ErrNoRecord
}
