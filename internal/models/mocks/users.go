package mocks

import (
	"time"

	"github.com/404GH0ST/snippetbox/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "gh0st@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "junken84@example.com" && password == "s3cur3p4ssw0rd" {
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
			Name:    "John",
			Email:   "junken84@example.com",
			Created: time.Now(),
		}

		return u, nil
	}
	return models.User{}, models.ErrNoRecord
}
