package services

import (
	"errors"
	"hexagonal/internals/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(email string, password string, confirmPass string) error {
	if password != confirmPass {
		return errors.New("the passwords are not equal")
	}
	err := s.userRepository.Register(email, password)
	if err != nil {
		return err
	}
	return nil
}

var _ ports.UserService = (*UserService)(nil)
