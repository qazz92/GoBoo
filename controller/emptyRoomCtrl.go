package controller

import (
	"github.com/qazz92/goTest/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"github.com/qazz92/goTest/redis"
)

type emptyRoom struct {
	Room_no string
}

func GetEmptyRoom(day int, from int, to int) []emptyRoom {
	key := strings.Join([]string{"empty_", strconv.Itoa(day), "_", strconv.Itoa(from),"_",strconv.Itoa(to)}, "")

	redisResult := redis.GetValueFromRedis(key)

	if len(redisResult) != 0 {

		fmt.Println("GetEmptyRoom cache!")


		var redisResultToEmptySlice []emptyRoom

		err :=json.Unmarshal([]byte(redisResult),&redisResultToEmptySlice)

		if err != nil {
			fmt.Println("json Err")
		}
		return redisResultToEmptySlice

	} else {

		fmt.Println("GetEmptyRoom db")

		db := mysql.GetDBCon()

		defer db.Close()

		rows, err := db.Query("SELECT b.room_no as room_no FROM timeTables a LEFT JOIN rooms b ON a.room_id = b.id WHERE a.day=? AND (a.time BETWEEN ? and ?) AND a.subject_code='빈 강의실' GROUP BY a.room_id HAVING count(*)=?",day,from,to,(to-from)+1 )
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close() //반드시 닫는다 (지연하여 닫기)


		var roomSlice []emptyRoom

		var room emptyRoom

		for rows.Next() {
			err := rows.Scan(&room.Room_no)
			if err != nil {
				log.Fatal(err)
			}
			roomSlice = append(roomSlice,room)
		}

		roomSliceMarshal,err := json.Marshal(roomSlice)

		redis.SetValueToRedis(key,roomSliceMarshal,0)

		return roomSlice
	}
}
