package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"time"
)

func (r *repository) SetString(key, data string, ttl time.Duration) error {
	err := r.client.Set(context.Background(), key, data, ttl).Err()
	if err != nil {
		utils.SubmitSentryLog(logSection, "SetString", err)
		return err
	}
	return nil
}

func (r *repository) GetString(key string) (string, error) {
	data, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		utils.SubmitSentryLog(logSection, "GetString", err)
		return "", err
	} else {
		return data, nil
	}
}

func (r *repository) Remove(key string) error {
	err := r.client.Del(context.Background(), key).Err()
	if err != nil {
		utils.SubmitSentryLog(logSection, "Remove", err)
		return derrors.InternalError()
	}
	return nil
}

func (r *repository) GetKeyTTl(key string) (time.Duration, error) {
	t, err := r.client.PTTL(context.Background(), key).Result()
	if err != nil {
		utils.SubmitSentryLog(logSection, "GetKeyTTl", err)
		return 0, derrors.InternalError()
	}
	return t, nil
}
