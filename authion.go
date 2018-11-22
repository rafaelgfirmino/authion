package main

import (
	"github.com/oftall/authion/infra/configuration"
	"github.com/oftall/authion/infra/server"
	"github.com/oftall/authion/infra/store"
)

func main() {
	configuration.Load()
	store.NewStore()
	server.Start()
}
