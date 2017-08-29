package controller

import (
	"log"
	"github.com/qazz92/GoBoo/mysql"
)

func GetPushPermit(stuId string) int {
	db := mysql.GetDBCon()

	defer db.Close()

	rows, err := db.Query("SELECT push_permit FROM normal_users WHERE id=?",stuId )
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	var pushPermit int

	for rows.Next() {
		err := rows.Scan(&pushPermit)
		if err != nil {
			log.Fatal(err)
		}
	}

	return pushPermit
}
