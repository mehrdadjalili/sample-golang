package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
)

func (h *Handler) ExistsEmail(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.ByEmail)
	if err := h.bindData(c, req, "exists-email", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.ExistsEmail(req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert.Status)
}

func (h *Handler) ExistsPhone(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.ByPhone)
	if err := h.bindData(c, req, "exists-phone", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.ExistsPhone(req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert.Status)
}

func (h *Handler) RegisterByPhone(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.RegisterByPhone)
	if err := h.bindData(c, req, "register-by-phone", lang...); err != nil {
		return err
	}
	c.Request().Header.Get("longitude")
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.RegisterByPhone(req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}

func (h *Handler) ReSendCode(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.ReSendCode)
	if err := h.bindData(c, req, "re-send-code", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.ReSendCode(req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}

func (h *Handler) ChangePassword(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.ChangePassword)
	if err := h.bindData(c, req, "change-password", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.ChangePassword(req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}

func (h *Handler) ForgotPassword(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.ForgotPassword)
	if err := h.bindData(c, req, "forgot-password", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.ForgotPassword(req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}

func (h *Handler) Login(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.Login)
	if err := h.bindData(c, req, "login", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.Login(req, c.RealIP())
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}

func (h *Handler) VerifyLogin(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.Verify)
	if err := h.bindData(c, req, "verify", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.VerifyLogin(req, c.RealIP())
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}

func (h *Handler) VerifyRegister(c echo.Context) error {
	lang := getLanguage(c)
	req := new(request.Verify)
	if err := h.bindData(c, req, "verify", lang...); err != nil {
		return err
	}
	err := req.Validate()
	if err != nil {
		return h.httpInvalidInputError(c, err)
	}
	insert, err := h.authService.VerifyRegister(req)
	if err != nil {
		return h.httpError(err, c)
	}
	return h.httpResponse(c, insert)
}
