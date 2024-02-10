package cmd

import (
	"fmt"
	redisRepository "github.com/mehrdadjalili/facegram_auth_service/src/repository/redis"
	sessionRepository "github.com/mehrdadjalili/facegram_auth_service/src/repository/session"
	userRepository "github.com/mehrdadjalili/facegram_auth_service/src/repository/user"
	codeRepository "github.com/mehrdadjalili/facegram_auth_service/src/repository/verify_code"
	accountService "github.com/mehrdadjalili/facegram_auth_service/src/service/account"
	authService "github.com/mehrdadjalili/facegram_auth_service/src/service/auth"
	sessionService "github.com/mehrdadjalili/facegram_auth_service/src/service/session"
	userService "github.com/mehrdadjalili/facegram_auth_service/src/service/user"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/mehrdadjalili/facegram_auth_service/src/transport/grpc"
	"github.com/mehrdadjalili/facegram_auth_service/src/transport/http/echo"
	Encryption "github.com/mehrdadjalili/facegram_common/pkg/encryption"
	_ "github.com/mehrdadjalili/facegram_common/pkg/logger"
	"github.com/mehrdadjalili/facegram_common/pkg/translator/i18n"

	config "github.com/mehrdadjalili/facegram_auth_service/src/config"
	"github.com/urfave/cli/v2"
)

var serveCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  serve,
}

var logSection = "cmd-serve"

func serve(c *cli.Context) error {
	cfg := new(config.Config)
	if err := config.ReadYAML("resources/development/config.yaml", cfg); err != nil {
		return err
	}
	config.ReadEnv(cfg)

	encryption := Encryption.New(cfg.Encryption.Key, "", "")

	err := os.Setenv("CLIENT_AUTHENTICATION_KEY", cfg.Server.ClientAccessToken)
	if err != nil {
		return err
	}

	err = os.Setenv("MANAGER_AUTHENTICATION_KEY", cfg.Server.ManagerAccessToken)
	if err != nil {
		return err
	}

	translator, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}

	redisRepo, err := redisRepository.New(cfg.Database.Redis)
	if err != nil {
		return err
	}

	sessionRepo, err := sessionRepository.New(cfg.Database.MongoDb.Url, cfg.Database.MongoDb.Database)
	if err != nil {
		return err
	}

	userRepo, err := userRepository.New(cfg.Database.MongoDb)
	if err != nil {
		return err
	}

	codeRepo, err := codeRepository.New(cfg.Database.MongoDb.Url, cfg.Database.MongoDb.Database)
	if err != nil {
		return err
	}

	accountSrv := accountService.New(*cfg, userRepo, codeRepo, encryption)
	authSrv := authService.New(*cfg, userRepo, codeRepo, sessionRepo, encryption, redisRepo)
	sessionSrv := sessionService.New(*cfg, sessionRepo, encryption)
	userSrv := userService.New(*cfg, userRepo, encryption)

	handler := echo.NewHttpHandler(&echo.HandlerFields{
		Cfg:            *cfg,
		Translator:     translator,
		AccountService: accountSrv,
		AuthService:    authSrv,
		SessionService: sessionSrv,
		Encryption:     encryption,
	})

	httpServer := echo.NewHttpServer(handler)
	go func() {
		if err := httpServer.Start(cfg.Server.Port); err != nil {
			utils.SubmitSentryLog(logSection, "serve", err)
		}
	}()

	grpcClientServer := grpc.NewClient(authSrv, userSrv)
	go func() {
		if err := grpcClientServer.StartClient(cfg.Server.GrpcClientPort); err != nil {
			utils.SubmitSentryLog(logSection, "serve", err)
		}
	}()

	fmt.Println("\nsuccessfully start rpc(client) server on port ", cfg.Server.GrpcClientPort)

	grpcManagerServer := grpc.NewManager(userSrv, sessionSrv)
	go func() {
		if err := grpcManagerServer.StartManager(cfg.Server.GrpcManagerPort); err != nil {
			utils.SubmitSentryLog(logSection, "serve", err)
		}
	}()

	fmt.Println("\nsuccessfully start rpc(manager) server on port ", cfg.Server.GrpcManagerPort)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	return nil
}
