package store

import (
	"database/sql"
	"fmt"
	"github.com/oftall/authion/infra/configuration"
)

var Mysql *sql.DB

func NewStore() {
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		configuration.Env.GetString("database.username"),
		configuration.Env.GetString("database.password"),
		configuration.Env.GetString("database.host"),
		configuration.Env.GetString("database.port"),
		configuration.Env.GetString("database.Name"),
	)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	fmt.Println(dataSource)
	//log.Fatal(db.Ping())
	Mysql = db
}
