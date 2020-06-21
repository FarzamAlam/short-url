package models

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var rdb *redis.Client

func init() {
	_ = godotenv.Load()
	host := os.Getenv("cacheDB_host")
	port := os.Getenv("cacheDB_port")
	uri := fmt.Sprintf("%s:%s", host, port)

	rdb = redis.NewClient(&redis.Options{
		Addr:     uri,
		DB:       0,
		Password: "",
	})
	pong, _ := rdb.Ping(ctx).Result()
	fmt.Println("Ping - ", pong)
	url := &URL{
		ShortURL:    "1e4c36",
		OriginalURL: "https://github.com/farzamalam",
	}
	_ = url.SetCacheURL()
	u, _ := url.GetCacheURL()
	fmt.Println("u : ", u)
}

func GetRedisClient() *redis.Client {
	return rdb
}

func (url *URL) SetCacheURL() error {
	return GetRedisClient().Set(ctx, url.ShortURL, url.OriginalURL, time.Hour*2).Err()
}

func (url *URL) GetCacheURL() (string, error) {
	return GetRedisClient().Get(ctx, url.ShortURL).Result()
}
