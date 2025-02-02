package config

import "github.com/rushairer/sso/utils"

var ServerPort = utils.GetEnv(
	"SERVER_PORT",
	"8080",
)
