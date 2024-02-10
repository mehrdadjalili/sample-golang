package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
)

func (h *Handler) SessionList(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.BasicList)
	if err := h.bindData(c, req, "basic-list", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.sessionService.List(user.User.ID, user.Session.Id, req.Page, req.PerPage)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert.Data)
}

func (h *Handler) DeleteSession(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.ById)
	if err := h.bindData(c, req, "by-id", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	res, err := h.sessionService.Delete(&request.ById{Id: req.Id}, user.Session.Id, user.User.ID)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, res.Message)
}
