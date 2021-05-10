package repository

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	WireSet = wire.NewSet(NewRepository)
)

type UserRepository interface {
	Authenticate(username string, password string) bool
}

type PgUserRepository struct {
	db *gorm.DB
}

func NewRepository(db gorm.DB) UserRepository {
	return &PgUserRepository{&db}
}

func (m *PgUserRepository) Authenticate(username string, password string) bool {
	user := m.find(username)
	return user.Password == password // TODO: hash password
}

func (m *PgUserRepository) find(username string) *AppUser {
	var appUser AppUser
	result := m.db.Where("username = ?", username).Find(&appUser)
	if result.Error != nil {
		panic(result.Error)
	}

	return &appUser
}
