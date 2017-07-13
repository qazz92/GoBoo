package controller

import (
	"runtime"
	"github.com/PuerkitoBio/goquery"
	"github.com/qazz92/goTest/redis"
	"fmt"
	"encoding/json"
	"sync"
	"github.com/qazz92/goTest/crawler"
)

type meal struct {
	Inter    string
	Bumin_kyo string
	Gang string
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
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
		return meal{Inter:redisResultToMeal.Inter,Bumin_kyo:redisResultToMeal.Bumin_kyo,Gang:redisResultToMeal.Gang}

	} else {

		fmt.Println("meal crawler!")

		runtime.GOMAXPROCS(MaxParallelism())

		var wg sync.WaitGroup

		total := getMealSelect(date)

		results := make([]string, 3)

		for idx := 7;  idx<=9; idx++ {
			wg.Add(1)

			go func(total *goquery.Selection, idx int) {
				defer wg.Done()
				result, _ := total.Eq(idx).Html()
				results[idx-7] = result
			}(total,idx)
		}
		wg.Wait()

		var crawlerResultToMeal meal

		crawlerResultToMeal.Inter = results[0]
		crawlerResultToMeal.Bumin_kyo = results[1]
		crawlerResultToMeal.Gang = results[2]

		crawlerResultToMealMarshal, err := json.Marshal(crawlerResultToMeal)

		if err != nil {
			fmt.Println("marshal error!")
		}

		redis.SetValueToRedis("meal_"+date,crawlerResultToMealMarshal,0)


		return crawlerResultToMeal
	}
}
