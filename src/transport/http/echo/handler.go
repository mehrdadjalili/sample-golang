package echo

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/service"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"log"
	"net/http"

	"github.com/mehrdadjalili/facegram_auth_service/resources/messages"

	"github.com/mehrdadjalili/facegram_common/pkg/encryption"

	echo "github.com/labstack/echo/v4"
	"github.com/mehrdadjalili/facegram_auth_service/src/config"
	"github.com/mehrdadjalili/facegram_common/pkg/translator"
)

type (
	Handler struct {
		cfg            config.Config
		translator     translator.Translator
		encryption     encryption.Encryption
		accountService service.Account
		authService    service.Auth
		sessionService service.Session
	}

	HandlerFields struct {
		Cfg            config.Config
		Translator     translator.Translator
		AccountService service.Account
		AuthService    service.Auth
		SessionService service.Session
		Encryption     encryption.Encryption
	}
)

func NewHttpHandler(h *HandlerFields) *Handler {
	if h.Translator == nil {
		log.Fatal("handler translator is nil")
	}

	if h.AccountService == nil {
		log.Fatal("handler account service is nil")
	}

	if h.AuthService == nil {
		log.Fatal("handler auth service is nil")
	}

	if h.SessionService == nil {
		log.Fatal("handler session service is nil")
	}

	return &Handler{
		cfg:            h.Cfg,
		translator:     h.Translator,
		accountService: h.AccountService,
		authService:    h.AuthService,
		sessionService: h.SessionService,
		encryption:     h.Encryption,
	}
}

func (h *Handler) bindData(c echo.Context, req interface{}, parent string, lang ...translator.Language) error {
	if err := c.Bind(req); err != nil {
		utils.SubmitSentryLog("echo", "bindData-"+parent, err)
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}
	if err := c.Validate(req); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}
	return nil
}
