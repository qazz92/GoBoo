package controller

import (
	"github.com/NaySoftware/go-fcm"
	"fmt"
	"log"
	"github.com/qazz92/GoBoo/mysql"
)

const (
	serverKey = "AAAAF-Mg1Us:APA91bEPLiM2psyQh2rzCyHUmiKjksmdg8wU-05nPcZeKfthwTpfkJ835-T18g8q5swJb19L9GGiTubEplU6L3fMPBzTZDEJTLEwoS7BiQN8no9iGi_HWPrnRWdq1QwBZzNkcN_Be5vE"
)

func Run() {
	fmt.Println("start!")
	db := mysql.GetDBCon()

	defer db.Close()


	rows, err := db.Query("select a.push_service_id FROM devices a JOIN normal_users b ON a.user_id = b.id JOIN user_circles c ON b.id = c.user_id WHERE c.circle_id=9")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() //반드시 닫는다 (지연하여 닫기)


	var token string
	var tokens []string

	for rows.Next() {
		err := rows.Scan(&token)
		if err != nil {
			log.Fatal(err)
		}
		tokens = append(tokens,token)
	}

	notification := fcm.NotificationPayload{}

	notification.Title = "BOO"

	notification.Body = "TEST"

	notification.Badge = "1"


	data := map[string]string{
		"title": "Hello World1",
		"body": "Happy Day",
		"content":"content",
	}

	c := fcm.NewFcmClient(serverKey)
	c.SetNotificationPayload(&notification)
	c.SetContentAvailable(true)
	c.SetPriority(fcm.Priority_HIGH)
	c.NewFcmRegIdsMsg(tokens,data)
	//c.AppendDevices(tokens)
	//c.AppendDevices(xds)

	status, err := c.Send()


	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}

}
