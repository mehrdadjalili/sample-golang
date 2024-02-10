package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
)

func (h *Handler) AccountProfile(c echo.Context) error {
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}

	insert, err := h.accountService.Profile(&request.ById{
		Id: user.User.ID,
	})
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}

func (h *Handler) EditAccount(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.EditAccount)
	if err := h.bindData(c, req, "edit-account", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	res, err := h.accountService.Edit(user.User.ID, req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, res.Message)
}

func (h *Handler) SetEmail(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.ByEmail)
	if err := h.bindData(c, req, "by-email", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.accountService.SetEmail(user.User.ID, req.Email)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert.Message)
}

func (h *Handler) SetPhone(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.ByPhone)
	if err := h.bindData(c, req, "by-phone", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	res, err := h.accountService.SetPhone(user.User.ID, req.Phone)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, res.Message)
}

func (h *Handler) VerifyEmail(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.VerifyChange)
	if err := h.bindData(c, req, "verify-change", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	res, err := h.accountService.VerifyEmail(user.User.ID, req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, res.Message)
}

func (h *Handler) VerifyPhone(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.VerifyChange)
	if err := h.bindData(c, req, "verify-change", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.accountService.VerifyPhone(user.User.ID, req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert.Message)
}

func (h *Handler) AccountReSendCode(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.ReSendCode)
	if err := h.bindData(c, req, "re-send-code", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	res, err := h.accountService.ReSendCode(user.User.ID, req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, res.Message)
}

func (h *Handler) ChangeAccountPassword(c echo.Context) error {
	lang := getLanguage(c)
	user := getUserInfo(c)
	if user == nil {
		return h.httpError(derrors.New(derrors.StatusUnauthorized, messages.Unauthorized), c)
	}
	req := new(request.ChangeAccountPassword)
	if err := h.bindData(c, req, "change-account-password", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.accountService.ChangePassword(user.User.ID, req.Old, req.New)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert.Message)
}
