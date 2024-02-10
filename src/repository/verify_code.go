package repository

import "github.com/mehrdadjalili/facegram_auth_service/src/entity/models"

type (
	VerifyCode interface {
		Create(data *models.VerifyCode) error
		ById(id string) (*models.VerifyCode, error)
		Edit(data *models.VerifyCode) error
		DeleteById(id string) error
		DeleteToDate(date int64) error
	}
)
