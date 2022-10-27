package Redis

import "github.com/go-redis/redis/v9"

func InitClient() *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return redisDB
}
