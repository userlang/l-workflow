package main

import (
	"workflow/naocs"
	"workflow/redis"
	"workflow/route"
)

func main() {
	naocs.Start()
	redis.Start()
	route.Start()

}
