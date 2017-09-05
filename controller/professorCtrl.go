package controller

import (
	"github.com/qazz92/GoBoo/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"encoding/json"
	"github.com/qazz92/GoBoo/redis"
)

type professor struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Room string `json:"room"`
	Tel string `json:"tel"`
	Email string `json:"email"`
	Major string `json:"major"`
	Pos string `json:"pos"`
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

		rows, err := db.Query("SELECT * FROM professors WHERE major = ?  ORDER BY name", major)
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
