package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/sso/config"
	"github.com/rushairer/sso/frontend/api/bootstrap"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("api crashed, error:", err)
		}
	}()

	log.Println("starting...")
	server := gin.Default()
	bootstrap.SetupServer(server)

	log.Println("running...")
	err := server.Run(fmt.Sprintf(":%s", config.ApiServerPort))
	if err != nil {
		log.Println(err)
	}
}
