package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/sso/config"
	"github.com/rushairer/sso/frontend/web/bootstrap"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("web crashed, error:", err)
		}
	}()

	log.Println("starting...")
	server := gin.Default()
	bootstrap.SetupServer(server)

	log.Println("running...")
	if gin.IsDebugging() {
		err := server.RunTLS(":443", "./frontend/web/resources/dev.cert.pem", "./frontend/web/resources/dev.key.pem")
		if err != nil {
			log.Println(err)
		}
	} else {
		err := server.Run(fmt.Sprintf(":%s", config.ServerPort))
		if err != nil {
			log.Println(err)
		}
	}
}
