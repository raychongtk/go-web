package repository

import (
	"errors"
	"github.com/google/uuid"
)

type UserRepository interface {
	Authenticate(username string, password string) (bool, error)
}

type MockRepository struct {
}

func NewRepository() UserRepository {
	return &MockRepository{}
}

func (m *MockRepository) Authenticate(username string, password string) (bool, error) {
	user, err := find(username)
	if err != nil {
		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	if !user.Status {
		return false, nil
	}

	return true, nil
}

func find(username string) (user *User, err error) {
	users := mockUsers()
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func mockUsers() []*User {
	users := make([]*User, 0)
	users = append(users, &User{
		ID:        uuid.New().String(),
		Username:  "test",
		Password:  "test",
		FirstName: "Andy",
		LastName:  "Lau",
		Status:    false,
	})

	users = append(users, &User{
		ID:        uuid.New().String(),
		Username:  "test1",
		Password:  "test1",
		FirstName: "Ray",
		LastName:  "Chan",
		Status:    true,
	})

	users = append(users, &User{
		ID:        uuid.New().String(),
		Username:  "test2",
		Password:  "test2",
		FirstName: "Jacky",
		LastName:  "Wong",
		Status:    true,
	})

	return users
}
