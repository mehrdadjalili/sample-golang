package repository

import "time"

type (
	Redis interface {
		SetString(key, data string, ttl time.Duration) error
		GetString(key string) (string, error)
		Remove(key string) error
		GetKeyTTl(key string) (time.Duration, error)
	}
)
