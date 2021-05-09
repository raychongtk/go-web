package service

import "github.com/raychongtk/go-web/repository"

type Service struct {
	userRepo repository.UserRepository
}

func ProvideService(repository repository.UserRepository) (*Service, error) {
	return &Service{userRepo: repository}, nil
}
