package config

import "github.com/rushairer/sso/utils"

var WebServerPort = utils.GetEnv(
	"SERVER_PORT",
	"8080",
)

var ApiServerPort = utils.GetEnv(
	"SERVER_PORT",
	"8081",
)
