package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/magiconair/properties"
)

var db *sql.DB

var PropertyFile = []string{"database.properties"}
var P, _ = properties.LoadFiles(PropertyFile, properties.UTF8, true)

func GetDB() *sql.DB {
	var err error

	host, _ := P.Get("database.host")
	port, _ := P.Get("database.port")
	user, _ := P.Get("database.user")
	pwd, _ := P.Get("database.password")
	dbname, _ := P.Get("database.dbname")
	sslmode, _ := P.Get("database.sslmode")

	if db == nil {
		connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + pwd + " dbname=" + dbname + " sslmode=" + sslmode

		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}
	return db
}
