package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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
	var err error

	db, err := sql.Open("mysql", config.MySQLDSN)
	if err != nil {
		log.Println("db error:", err)
	}

	if gin.IsDebugging() {
		migrations(db)
	}

	log.Println("starting...")
	server := gin.Default()
	bootstrap.SetupServer(server)

	log.Println("running...")
	err = server.Run(fmt.Sprintf(":%s", config.WebServerPort))
	if err != nil {
		log.Println(err)
	}
}

func migrations(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("mysql error:", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		config.MigrationsPath,
		"mysql",
		driver,
	)
	if err != nil {
		log.Println("migrations error:", err)
	}
	if err = m.Up(); err != nil {
		log.Println("migrations up error:", err)
	}
}
