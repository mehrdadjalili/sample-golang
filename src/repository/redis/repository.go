package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/mehrdadjalili/facegram_auth_service/src/config"
	repo "github.com/mehrdadjalili/facegram_auth_service/src/repository"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
)

type repository struct {
	client *redis.Client
}

var logSection = "repository->redis"

func New(cfg config.Redis) (repo.Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Server,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	ping := rdb.Ping(context.Background())
	if ping.Val() != "PONG" {
		err := errors.New("cannot create connection to redis server")
		utils.SubmitSentryLog(logSection, "DeleteToDate", err)
		return nil, err
	}
	return &repository{
		client: rdb,
	}, nil
}
