package repository

import (
	"errors"
	"github.com/google/wire"
	"github.com/raychongtk/go-web/util"
	"gorm.io/gorm"
)

var (
	WireSet = wire.NewSet(NewRepository)
)

type UserRepository interface {
	Authenticate(username string, password string) (bool, error)
}

type PgUserRepository struct {
	db *gorm.DB
}

func NewRepository(db gorm.DB) UserRepository {
	return &PgUserRepository{&db}
}

func (m *PgUserRepository) Authenticate(username string, password string) (bool, error) {
	user, err := m.find(username)
	if err != nil {
		return false, err
	}
	return util.ValidateHash(user.Password, password), nil
}

func (m *PgUserRepository) find(username string) (*AppUser, error) {
	var appUser AppUser
	result := m.db.Where("username = ?", username).Find(&appUser)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &appUser, nil
}
