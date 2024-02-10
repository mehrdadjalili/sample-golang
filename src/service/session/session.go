package session

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
)

func (s *srv) List(userId, currentSessionId string, page, perPage int) (*response.SessionList, error) {
	sessions, err := s.sessionRepo.UserSessions(userId, page, perPage)
	if err != nil {
		return nil, err
	}

	var lst []response.Session

	for _, item := range sessions {
		var thisSession = false
		if currentSessionId == item.Id.Hex() {
			thisSession = true
		}
		lst = append(lst, response.Session{
			Id:          item.Id.Hex(),
			DeviceName:  item.DeviceName,
			CreatedAt:   item.CreatedAt,
			ThisSession: thisSession,
		})
	}

	return &response.SessionList{
		Data: lst,
	}, nil
}

func (s *srv) Delete(req *request.ById, currentSessionId, userId string) (*response.Regular, error) {
	if req.Id == currentSessionId {
		return nil, derrors.New(derrors.StatusBadRequest, messages.BadRequest)
	}

	err := s.sessionRepo.DeleteById(req.Id, userId)
	if err != nil {
		return nil, err
	}

	return &response.Regular{Message: "ok"}, nil
}
