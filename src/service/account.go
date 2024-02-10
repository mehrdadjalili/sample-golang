package service

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
)

type (
	Account interface {
		Profile(req *request.ById) (*response.Profile, error)
		Edit(userId string, req *request.EditAccount) (*response.Regular, error)
		SetEmail(userId string, email string) (*response.Regular, error)
		SetPhone(userId string, phone string) (*response.Regular, error)
		VerifyEmail(userId string, req *request.VerifyChange) (*response.Regular, error)
		VerifyPhone(userId string, req *request.VerifyChange) (*response.Regular, error)
		ReSendCode(userId string, req *request.ReSendCode) (*response.Regular, error)
		ChangePassword(userId string, old, new string) (*response.Regular, error)
	}
)
