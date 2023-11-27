package usecase

import (
	"fmt"
	"go_laundry/model/dto"
	"go_laundry/repository"
	"go_laundry/util/security"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequest) (dto.AuthRespone, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(payload dto.AuthRequest) (dto.AuthRespone, error) {
	//username di database
	user, err := a.userRepo.FindByUsername(payload.Username)
	if err != nil {
		return dto.AuthRespone{}, fmt.Errorf("unauthorized: invalid credential")
	}
	//validasi password
	err = security.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		return dto.AuthRespone{}, fmt.Errorf("unauthorized: invalid credential")
	}
	//Generate Token
	token, err := security.GenerateJWTToken(user)
	if err != nil {
		return dto.AuthRespone{}, err
	}
	return dto.AuthRespone{
		Username: user.Username,
		Token:    token,
	}, nil
}

func NewAuthUseCase(userRepo repository.UserRepository) AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
	}
}
