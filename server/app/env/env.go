package env

import (
	"go.uber.org/zap"
)

type Env struct {
	Logger *zap.SugaredLogger
	Cache *Cache
}

func NewEnv() (*Env, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	logger := l.Sugar()

	cache, err := NewCacheClient()
	if err != nil {
		return nil, err
	}

	return &Env{Logger: logger, Cache: cache}, nil
}
