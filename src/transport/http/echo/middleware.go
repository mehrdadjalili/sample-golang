package echo

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
	"strings"
)

func (h *Handler) authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := getToken(c)
		if token == "" {
			return &echo.HTTPError{Code: 401, Message: "Unauthorized"}
		}
		checkToken, err := h.authService.CheckToken(token)
		if err != nil {
			if err.Error() == "Unauthorized" {
				return &echo.HTTPError{Code: 401, Message: "Unauthorized"}
			}
			message, code := derrors.HttpError(err)
			lang := getLanguage(c)
			return &echo.HTTPError{
				Code:    code,
				Message: h.translator.Translate(message, lang...),
			}
		}

		js, err := json.Marshal(checkToken)
		if err != nil {
			message, code := derrors.HttpError(err)
			lang := getLanguage(c)
			return &echo.HTTPError{
				Code:    code,
				Message: h.translator.Translate(message, lang...),
			}
		}

		c.Set("userInfo", string(js))

		return next(c)
	}
}

func getToken(c echo.Context) string {

	header := c.Request().Header
	authv := header.Get("Authorization")

	if !strings.HasPrefix(strings.ToLower(authv), "bearer") {
		return ""
	}

	values := strings.Split(authv, " ")
	if len(values) < 2 {
		return ""
	}

	return values[1]
}

func (h *Handler) getUserIdFromToken(c echo.Context) (string, error) {
	token := getToken(c)
	data, err := utils.ParseJwtToken(h.cfg.Encryption.Key, token)
	if err != nil {
		return "", err
	}
	dec, err := h.encryption.AesDecrypt(data)

	var userIfo map[string]interface{}

	err = json.Unmarshal([]byte(dec), &userIfo)
	if err != nil {
		return "", err
	}

	userId, existsUserId := userIfo["user_id"]

	if !existsUserId {
		return "", derrors.New(derrors.StatusUnauthorized, messages.Unauthorized)
	}

	return userId.(string), nil
}
