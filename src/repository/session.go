package repository

import "github.com/mehrdadjalili/facegram_auth_service/src/entity/models"

type (
	Session interface {
		Create(data *models.Session) error
		ById(id string) (*models.Session, error)
		Edit(data *models.Session) error
		DeleteById(id, userId string) error
		DeleteByIds(ids []string) error
		DeleteToDate(date int64) error
		UserSessions(userId string, page, limit int) ([]models.Session, error)
	}
)
