package service

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
)

type (
	Session interface {
		List(userId, currentSessionId string, page, perPage int) (*response.SessionList, error)
		Delete(req *request.ById, currentSessionId, userId string) (*response.Regular, error)
	}
)
