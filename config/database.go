package config

import "github.com/rushairer/sso/utils"

var MySQLDSN = utils.GetEnv(
	"MYSQL_DSN",
	"root:123456@tcp(localhost:3306)/sso?multiStatements=true&parseTime=true",
)

var MigrationsPath = utils.GetEnv(
	"MIGRATIONS_PATH",
	"file://./migrations",
)

var RedisDSN string = utils.GetEnv(
	"REDIS_DSN",
	"127.0.0.1:6379",
)
