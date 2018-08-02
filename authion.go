package main

import (
	"github.com/rafaelgfirmino/authion/infra/configuration"
	"github.com/rafaelgfirmino/authion/infra/server"
	)

func main() {
	configuration.Load()
	server.Start()
}
