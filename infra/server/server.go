package server

import (
	"fmt"
	"github.com/rafaelgfirmino/authion/infra/configuration"
	"github.com/rafaelgfirmino/authion/infra/router"
	"log"
	"net/http"
)

//Initialize server HTTP
func Start() {
	port := fmt.Sprintf(":%d", configuration.Env.GetInt("server.port"))
	log.Fatal(http.ListenAndServe(port, router.Router))

}
