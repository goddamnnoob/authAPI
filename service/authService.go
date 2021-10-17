package service

import (
	"github.com/goddamnnoob/authAPI/dto"
	"github.com/goddamnnoob/notReddit/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}
