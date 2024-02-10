package session

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/config"
	"github.com/mehrdadjalili/facegram_auth_service/src/repository"
	"github.com/mehrdadjalili/facegram_auth_service/src/service"
	"github.com/mehrdadjalili/facegram_common/pkg/encryption"
)

type (
	srv struct {
		cfg         config.Config
		sessionRepo repository.Session
		encryption  encryption.Encryption
	}
)

func New(cfg config.Config, sessionRepo repository.Session,
	encryption encryption.Encryption) service.Session {
	return &srv{
		cfg:         cfg,
		sessionRepo: sessionRepo,
		encryption:  encryption,
	}
}
