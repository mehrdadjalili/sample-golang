package repository

import "github.com/mehrdadjalili/facegram_auth_service/src/entity/models"

type (
	User interface {
		Create(data *models.User) error
		ById(id string) (*models.User, error)
		Edit(data *models.User) error
		ExistsEmail(hashedEmail string) (bool, error)
		ExistsPhone(hashedPhone string) (bool, error)
		ByEmailOrPhone(user string) (*models.User, error)
		CountByRole(isAgent bool) (int64, error)
		Count() (int64, error)
		List(search, sort string, page, perPage int) ([]models.User, int64, error)
	}
)
