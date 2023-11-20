package main

import (
	"workflow/redis"
	"workflow/route"
)

func main() {
	//  naocs.Start()
	redis.Start()
	route.Start()

}
