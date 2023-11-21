package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"workflow/config"
)

func Start() {

	s := connection()
	if s != "" {

		fmt.Println("链接到的redis", s)

	}

}
func connection() string {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddressPort,
		Password: config.RedisPwd,
		DB:       config.RedisDB,
	})

	result, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}
