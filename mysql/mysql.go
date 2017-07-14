package mysql

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetDBCon()  *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("INMISHOST")
	db, err := sql.Open("mysql", "boo_npe:ekfqlc152!@tcp("+url+":3306)/boo")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}


	return db
}
