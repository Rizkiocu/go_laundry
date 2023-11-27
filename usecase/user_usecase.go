package usecase

import (
	"fmt"
	"go_laundry/model"
	"go_laundry/model/dto"
	"go_laundry/repository"
	"go_laundry/util/common"
	"go_laundry/util/security"
)

type UserUseCase interface {
	FindByUsername(username string) (model.UserCredential, error)
	Register(payload dto.AuthRequest) error
}
type userUseCase struct {
	userRepo repository.UserRepository
}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByUsername(username string) (model.UserCredential, error) {
	return u.userRepo.FindByUsername(username)
}

// Register implements UserUseCase.
func (u *userUseCase) Register(payload dto.AuthRequest) error {
	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	userCredential := model.UserCredential{
		Id:       common.GenerateID(),
		Username: payload.Username,
		Password: hashPassword,
	}
	err = u.userRepo.Save(userCredential)
	if err != nil {
		return fmt.Errorf("failde user Save: %v", err.Error())
	}
	return nil
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
