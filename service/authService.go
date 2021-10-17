package service

import (
	"github.com/goddamnnoob/authAPI/domain"
	"github.com/goddamnnoob/authAPI/dto"
	"github.com/goddamnnoob/notReddit/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	var appErr *errs.AppError
	var login *domain.Login
	login, appErr := s.repo.FindBy(req.Username, req.Password)
	if appErr != nil {
		return nil, appErr
	}
	return &dto.LoginResponse{}, appErr
}
