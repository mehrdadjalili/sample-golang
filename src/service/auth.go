package service

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
)

type (
	Auth interface {
		ExistsEmail(req *request.ByEmail) (*response.Status, error)
		ExistsPhone(req *request.ByPhone) (*response.Status, error)
		RegisterByPhone(req *request.RegisterByPhone) (*response.ResultData, error)
		ReSendCode(req *request.ReSendCode) (*response.ResultData, error)
		ChangePassword(req *request.ChangePassword) (*response.ResultData, error)
		ForgotPassword(req *request.ForgotPassword) (*response.ResultData, error)
		Login(req *request.Login, ip string) (*response.ResultData, error)
		VerifyLogin(req *request.Verify, ip string) (*response.ResultData, error)
		VerifyRegister(req *request.Verify) (*response.ResultData, error)
		CheckToken(token string) (*response.CheckToken, error)
	}
)
