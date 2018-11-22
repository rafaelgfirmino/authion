package store

import (
	"database/sql"
	"fmt"
	"github.com/oftall/authion/infra/configuration"
	"log"
	"os"
)

var Mysql *sql.DB

func NewStore() {
	logger := log.New(os.Stdout, "database ", log.LstdFlags)

	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		configuration.Env.GetString("database.user"),
		configuration.Env.GetString("database.password"),
		configuration.Env.GetString("database.host"),
		configuration.Env.GetString("database.port"),
		configuration.Env.GetString("database.name"),
	)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Database error conection: \n",err))
	}
	logger.Println("Database is connected")
	Mysql = db
}
