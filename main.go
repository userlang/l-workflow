package main

import (
	"workflow/naocs"
	"workflow/redis"
	"workflow/repository"
	"workflow/route"
)

func main() {
	naocs.Start()
	repository.InitDB()
	redis.Start()
	route.Start()

}
