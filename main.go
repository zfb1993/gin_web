package main

import (
	"gin_web/routers"
)

func main()  {
	router := routers.InitRouter()
	router.Static("/static","./static")
	router.Run(":7777")
}