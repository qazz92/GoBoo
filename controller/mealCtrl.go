package controller

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"encoding/json"
	"github.com/qazz92/GoBoo/crawler"
	"github.com/qazz92/GoBoo/redis"
)

type meal struct {
	Inter    string `json:"inter"`
	Bumin_kyo string `json:"bumin_kyo"`
	Gang string `json:"gang"`
	Hadan_kyo string `json:"hadan_kyo"`
	Hadan_gang string `json:"hadan_gang"`
	Library string `json:"library"`
}

func getMealSelect(date string) *goquery.Selection {

	const url  =  "http://www.donga.ac.kr/MM_PAGE/SUB007/SUB_007005005.asp?PageCD=007005005&seldate="
	doc := crawler.GETDoc(url+date)
	total:= doc.Find("#subContext").Eq(0).Find("tr").Eq(1).Find(".sk01TBL").Eq(1).Find(".sk03TBL").Find("td")
	return total
}

func GetMeal(date string) meal{

	redisResult := redis.GetValueFromRedis("meal_"+date)

	if len(redisResult) != 0 {

		fmt.Println("meal cache!")

		var redisResultToMeal meal

		err :=json.Unmarshal([]byte(redisResult),&redisResultToMeal)

		if err != nil {
			fmt.Println("json Err")
		}
		return meal{Inter:redisResultToMeal.Inter,Bumin_kyo:redisResultToMeal.Bumin_kyo,Gang:redisResultToMeal.Gang,Hadan_kyo:redisResultToMeal.Hadan_kyo,Hadan_gang:redisResultToMeal.Hadan_gang,Library:redisResultToMeal.Library}

	} else {

		fmt.Println("meal crawler!")

		total := getMealSelect(date)

		results := make([]string, 6)

		count := 0

		for idx := 0;  idx<=9; idx++ {
			if idx==0 || idx==1 || idx==3 || idx==7 || idx==8 || idx==9 {
				result, _ := total.Eq(idx).Html()
				results[count] = result
				count += 1

			}
		}
		var crawlerResultToMeal meal

		fmt.Println(len(results))

		crawlerResultToMeal.Hadan_kyo = results[0]
		crawlerResultToMeal.Hadan_gang = results[1]
		crawlerResultToMeal.Library = results[2]
		crawlerResultToMeal.Inter = results[3]
		crawlerResultToMeal.Bumin_kyo = results[4]
		crawlerResultToMeal.Gang = results[5]

		crawlerResultToMealMarshal, err := json.Marshal(crawlerResultToMeal)

		if err != nil {
			fmt.Println("marshal error!")
		}

		redis.SetValueToRedis("meal_"+date,crawlerResultToMealMarshal,0)

		return crawlerResultToMeal
	}
}
