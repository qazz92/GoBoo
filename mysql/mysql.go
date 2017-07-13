package mysql

import "database/sql"

func GetDBCon()  *sql.DB {
	db, err := sql.Open("mysql", "boo_npe:ekfqlc152!@tcp(168.115.128.42:3306)/boo")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}


	return db
}
