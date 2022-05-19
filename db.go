package godjan

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB
var variabel = "sdads"

func init() {
	if os.Getenv("POSTGRES_SSLMODE") == "" {
		os.Setenv("POSTGRES_SSLMODE", "disable")
	}
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_SSLMODE"),
	)
	db, _ = sql.Open("postgres", psqlconn)
}

func SetConnection(db *sql.DB) {
	db = db
}

func MakeMigrations(db *sql.DB) {

}

func GetParameter() {
	fmt.Println(variabel)
}

func SetParameter(arg string) {
	variabel = arg
}
