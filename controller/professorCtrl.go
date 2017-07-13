package controller

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/qazz92/goTest/redis"
	"fmt"
	"encoding/json"
	"github.com/qazz92/goTest/mysql"
)

type professor struct {
	Id int
	Name string
	Room string
	Tel string
	Email string
	Major string
	Pos string
}

func GetProfessors(major string) []professor {

	key := "professor_list_"+major

	redisResult := redis.GetValueFromRedis(key)

	if len(redisResult) != 0 {

		fmt.Println("professor cache!")

		var redisResultToProfessorSlice []professor

		err :=json.Unmarshal([]byte(redisResult),&redisResultToProfessorSlice)

		if err != nil {
			fmt.Println("json Err")
		}
		return redisResultToProfessorSlice


	} else {
		fmt.Println("professor db")

		db := mysql.GetDBCon()

		defer db.Close()

		rows, err := db.Query("SELECT * FROM professors WHERE major = ?", major)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close() //반드시 닫는다 (지연하여 닫기)

		var professorInMajorSclice []professor

		var p professor

		for rows.Next() {
			err := rows.Scan(&p.Id,&p.Name,&p.Room,&p.Tel,&p.Email,&p.Major,&p.Pos)
			if err != nil {
				log.Fatal(err)
			}
			professorInMajorSclice = append(professorInMajorSclice,p)
		}

		professorInMajorMarshal,err := json.Marshal(professorInMajorSclice)

		redis.SetValueToRedis(key,professorInMajorMarshal,0)

		return professorInMajorSclice
	}
}
