package service

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
)

type (
	User interface {
		ById(id string) (*response.User, error)
		Edit(data *request.User) (*response.Regular, error)
		CountByRole(isAgent bool) (int64, error)
		Count() (int64, error)
		List(search, sort string, page, perPage int) ([]response.User, int64, error)
	}
)
