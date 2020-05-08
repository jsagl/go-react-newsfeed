package env

import (
	"github.com/go-redis/redis/v7"
	"os"
	"time"
)

type Cache struct {
	Client *redis.Client
}

func NewCacheClient() (*Cache, error) {
	opt, e := redis.ParseURL(os.Getenv("REDISCLOUD_URL"))
	if e != nil {
		return nil, e
	}

	client := redis.NewClient(opt)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Cache{client}, nil
}


func (c *Cache) Get(key string) ([]byte, error) {
	cache, err := c.Client.Get(key).Bytes()

	if err != nil && err.Error() == "redis: nil" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return cache, nil
}

func (c *Cache) Set(key string, val interface{}, duration time.Duration) {
	c.Client.Set(key, val, duration)
}