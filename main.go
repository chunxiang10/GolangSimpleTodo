package main

import (
	"simple-todo/config"
	"simple-todo/router"
)

func main() {
	conf := config.Conf()
	r := router.InitRouter()
	r.Run(conf.Server.Host + ":" + conf.Server.Port)
}
