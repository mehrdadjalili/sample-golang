package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func rateConfig() middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			code := http.StatusForbidden
			return sendResponse(context, nil, code, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			code := http.StatusTooManyRequests
			return sendResponse(context, nil, code, nil)
		},
	}
}

func sendResponse(c echo.Context, msg *string, code int, result interface{}) error {
	data := map[string]interface{}{
		"message": msg,
		"code":    code,
		"data":    result,
	}
	return c.JSON(http.StatusOK, data)
}
