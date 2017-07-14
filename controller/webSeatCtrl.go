package controller

import (
	"strings"
	"regexp"
	"github.com/qazz92/GoBoo/crawler"
)

type seat struct {
	Loc string `json:"loc"`
	All string `json:"all"`
	Use string `json:"use"`
	Remain string `json:"remain"`
	Util string `json:"util"`
}

func GetWebSeat() []seat {

	const url = "http://168.115.33.207/WebSeat"

	doc := crawler.GETDoc(url)

	total := doc.Find("table").Eq(1)

	var webSeatSlice []seat
	for idx:=12; idx<23;idx++  {
		row := total.Find("tr").Eq(idx)
		var webSeat seat
		for rowIdx:=0; rowIdx<5;rowIdx++  {

			result := row.Find("td").Eq(rowIdx).Text()
			cleanResult := strings.Replace(result,"\xc2\xa0","",1)
			switch rowIdx {
			case 0:
				webSeat.Loc = cleanResult
				break
			case 1:
				webSeat.All = cleanResult
				break
			case 2:
				webSeat.Use = cleanResult
				break
			case 3:
				webSeat.Remain = cleanResult
				break
			case 4:
				re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$+|\s+\%`)
				webSeat.Util = re_leadclose_whtsp.ReplaceAllString(cleanResult, "")
				break
			}
		}
		webSeatSlice = append(webSeatSlice, webSeat)
	}

	return webSeatSlice
}

