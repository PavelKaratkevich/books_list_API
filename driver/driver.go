package driver

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var db *sqlx.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sqlx.DB {
	gotenv.Load()
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sqlx.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
