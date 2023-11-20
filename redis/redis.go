package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func Start() {

	s := connection()
	if s != "" {

		fmt.Println("链接到的redis", s)

	}

}
func connection() string {

	client := redis.NewClient(&redis.Options{
		Addr:     "172.16.0.166:22001",
		Password: "w3DEhqL67fF12@",
		DB:       0,
	})

	result, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}
