package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: r,
	}
}

func (u *UserUseCase) GetUserByID(ID int) (*model.User, error) {
	return u.repository.FindById(ID)
}

func (u *UserUseCase) CreateUser(user model.User) (*model.User, error) {
	return u.repository.CreateUser(user)
}
