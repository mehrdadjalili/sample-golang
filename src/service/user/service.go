package user

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/config"
	"github.com/mehrdadjalili/facegram_auth_service/src/repository"
	"github.com/mehrdadjalili/facegram_auth_service/src/service"
	"github.com/mehrdadjalili/facegram_common/pkg/encryption"
)

type (
	srv struct {
		cfg        config.Config
		userRepo   repository.User
		encryption encryption.Encryption
	}
)

func New(cfg config.Config, userRepo repository.User,
	encryption encryption.Encryption) service.User {
	return &srv{
		cfg:        cfg,
		userRepo:   userRepo,
		encryption: encryption,
	}
}
