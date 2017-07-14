package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/qazz92/GoBoo/controller"
	"runtime"
	"github.com/sclevine/agouti"

	"github.com/PuerkitoBio/goquery"

	"strings"

)
func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	runtime.GOMAXPROCS(MaxParallelism())
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/meal", func(c *gin.Context) {
		date := c.Query("date")
		c.JSON(http.StatusOK, gin.H{"result_code":http.StatusOK,"result_body":controller.GetMeal(date)})
	})
	r.GET("/getWebSeat", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result_code":http.StatusOK,"result_body":controller.GetWebSeat()})
	})
	r.GET("/getPro", func(c *gin.Context) {
		major := c.Query("major")
		c.JSON(http.StatusOK, gin.H{"result_code":http.StatusOK,"result_body":controller.GetProfessors(major)})
	})

	r.GET("/getEmptyRoom", func(c *gin.Context) {
		day, dayErr := strconv.Atoi(c.Query("day"))
		if dayErr != nil {
			panic("dayErr")
		}
		from, fromErr := strconv.Atoi(c.Query("from"))
		if fromErr != nil {
			panic("fromErr")
		}
		to, toErr := strconv.Atoi(c.Query("to"))
		if toErr != nil {
			panic("toErr")
		}
		c.JSON(http.StatusOK, gin.H{"result_code":http.StatusOK,"result_body":controller.GetEmptyRoom(day,from,to)})
	})

	r.GET("/login", func(c *gin.Context) {
		webDriver := agouti.ChromeDriver() // ChromeDriverを使う
		if err := webDriver.Start(); err != nil {
			panic(err)
		}
		defer webDriver.Stop()

		page, err := webDriver.NewPage()

		if err != nil {
			panic(err)
		}
		// Navigateで指定したURLにアクセスする
		if err := page.Navigate("https://student.donga.ac.kr/Univ/SUG/SSUG0020.aspx?m=3"); err != nil {
			panic(err)

		}
		if err := page.FindByID("txtStudentCd").Fill("1124305"); err != nil {
			panic(err)
		}
		if err := page.FindByID("txtPasswd").Fill("Ekfqlc152!"); err != nil {
			panic(err)
		}
		if err := page.FindByID("ibtnLogin").Click(); err != nil {
			panic(err)
		}
		if err := page.FindByID("ddlYear").Select("2017"); err != nil {
			panic(err)
		}
		if err := page.FindByID("ddlSmt").Select("1학기"); err != nil {
			panic(err)
		}
		if err := page.FindByID("ImageButton1").Click(); err != nil {
			panic(err)
		}
		//if err := page.Navigate("https://student.donga.ac.kr/Univ/SUD/SSUD0000.aspx?m=1"); err != nil {
		//	panic(err)
		//
		//}
		html,_ := page.HTML()
		html_code  := strings.NewReader(html)

		doc, _ := goquery.NewDocumentFromReader(html_code)

		var timeTemp []string

		doc.Find("#dgRep").Find("tr").Find("td").Each(func(i int, s *goquery.Selection) {

			if i>15 {
				timeTemp = append(timeTemp,strings.TrimSpace(s.Text()))
			}
		})

		var timeTable [][]string
		end := 14

		for start := 0; start < len(timeTemp); start+=15 {
			if end > len(timeTemp) { end=len(timeTemp) }
			temp := timeTemp[start:end]
			timeTable = append(timeTable,temp)
			end += 15
		}

		c.JSON(http.StatusOK,gin.H{"result_code":http.StatusOK,"result_body":timeTable})
	})

	r.Run(":3000")
}
