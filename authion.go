package main

import (
	"github.com/rafaelgfirmino/authion/infra/configuration"
	"github.com/rafaelgfirmino/authion/infra/server"
	"github.com/rafaelgfirmino/authion/infra/store"
)

func main() {
	configuration.Load()
	server.Start()
	store.NewStore()
}
