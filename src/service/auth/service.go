package auth

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
		sessionRepo    repository.Session
		encryption     encryption.Encryption
		redis          repository.Redis
	}
)

func New(cfg config.Config, userRepo repository.User,
	verifyCodeRepo repository.VerifyCode,
	sessionRepo repository.Session,
	encryption encryption.Encryption,
	redis repository.Redis,
) service.Auth {
	return &srv{
		cfg:            cfg,
		userRepo:       userRepo,
		verifyCodeRepo: verifyCodeRepo,
		sessionRepo:    sessionRepo,
		encryption:     encryption,
		redis:          redis,
	}
}
