package main

import (
	"workflow/naocs"
	"workflow/route"
)

func main() {
	naocs.Init()
	route.Start()

}
