package usecase

import (
	"go-api/model"
	"go-api/repository"
	"go-api/security"
)

type LoginUseCase struct {
	repository repository.UserRepository
}

func NewLoginUseCase(repository repository.UserRepository) LoginUseCase {
	return LoginUseCase{
		repository: repository,
	}
}

func (u *LoginUseCase) Login(credentials model.User) (*model.User, string, error) {
	user, err := u.repository.FindByName(credentials.Name)

	if err != nil {
		return nil, "", err
	}

	err = security.VerifyPassword(user.Password, credentials.Password)
	if err != nil {
		return nil, "", err
	}

	token, err := security.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
