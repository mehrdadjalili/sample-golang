package echo

import (
	"context"
	"fmt"
	"net/http"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mehrdadjalili/facegram_auth_service/src/transport"
)

type (
	httpServer struct {
		handler *Handler
		auth    *echo.Group
		account *echo.Group
		session *echo.Group
	}
)

var (
	e = echo.New()
)

func NewHttpServer(handler *Handler) transport.Echo {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Validator = &customValidator{Validator: validator.New()}

	e.Use(middleware.RateLimiterWithConfig(rateConfig()))

	public := e.Group("/v1")

	account := public.Group("/account")
	account.Use(handler.authenticate)

	session := public.Group("/session")
	session.Use(handler.authenticate)

	return &httpServer{
		handler: handler,
		auth:    public.Group("/auth"),
		account: account,
		session: session,
	}
}

func (s *httpServer) Start(port int) error {
	s.setRoutes()

	if port == 0 {
		port = 8080
	}

	return e.Start(fmt.Sprintf(":%d", port))
}

func (s *httpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // use config for time
	defer cancel()

	return e.Shutdown(ctx)
}
