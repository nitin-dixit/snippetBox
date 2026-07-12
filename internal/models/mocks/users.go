package mocks

import "github.com/nitin-dixit/snippetBox/internal/models"

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
