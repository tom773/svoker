package main

import (
	"github.com/tom773/svoker/api/routes"
	"github.com/tom773/svoker/api/ws"
)

func main() {
	go ws.InitWS()
	routes.SetupRoutes()
	select {}
}
