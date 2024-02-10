package echo

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"github.com/mehrdadjalili/facegram_common/pkg/translator"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type langQ struct {
	lang string
	q    float64
}

func getLanguage(c echo.Context) []translator.Language {
	acceptLanguages := c.Request().Header.Get("Accept-Language")

	var lqs []langQ
	languages := strings.Split(acceptLanguages, ",")

	for _, language := range languages {
		language = strings.Trim(language, " ")
		langWithQ := strings.Split(language, ";")

		if len(langWithQ) == 1 {
			lq := langQ{langWithQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			valueQ := strings.Split(langWithQ[1], "=")
			q, err := strconv.ParseFloat(valueQ[1], 64)
			if err != nil {
				continue
			}
			lq := langQ{langWithQ[0], q}
			lqs = append(lqs, lq)
		}
	}

	sort.SliceStable(lqs, func(i, j int) bool {
		return lqs[i].q > lqs[j].q
	})

	var result []translator.Language
	for _, lq := range lqs {
		result = append(result, translator.Language(lq.lang))
	}

	return result
}

func getUserInfo(c echo.Context) *response.CheckToken {
	u := c.Get("userInfo")
	if u == nil {
		return nil
	}
	var user response.CheckToken
	err := json.Unmarshal([]byte(u.(string)), &user)
	if err != nil {
		return nil
	}
	return &user
}

func (h *Handler) httpError(e error, c echo.Context) error {
	message, code := derrors.HttpError(e)
	data := map[string]interface{}{
		"message": message,
		"code":    code,
		"data":    nil,
	}
	return c.JSON(http.StatusUnprocessableEntity, data)
}

func (h *Handler) httpInvalidInputError(c echo.Context, e error) error {
	data := map[string]interface{}{
		"message": e.Error(),
		"code":    http.StatusUnprocessableEntity,
		"data":    nil,
	}
	return c.JSON(http.StatusUnprocessableEntity, data)
}

func (h *Handler) httpResponse(c echo.Context, result interface{}) error {
	data := map[string]interface{}{
		"message": nil,
		"code":    http.StatusOK,
		"data":    result,
	}
	return c.JSON(http.StatusOK, data)
}

func getUserIp(c echo.Context) string {
	return strings.Replace(c.RealIP(), ".", "-", 10)
}
