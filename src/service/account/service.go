package account

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/config"
	"github.com/mehrdadjalili/facegram_auth_service/src/repository"
	"github.com/mehrdadjalili/facegram_auth_service/src/service"
	"github.com/mehrdadjalili/facegram_common/pkg/encryption"
)

type (
	srv struct {
		cfg            config.Config
		userRepo       repository.User
		verifyCodeRepo repository.VerifyCode
		encryption     encryption.Encryption
	}
)

func New(cfg config.Config, userRepo repository.User,
	verifyCodeRepo repository.VerifyCode,
	encryption encryption.Encryption) service.Account {
	return &srv{
		encryption:     encryption,
		userRepo:       userRepo,
		verifyCodeRepo: verifyCodeRepo,
		cfg:            cfg,
	}
}
