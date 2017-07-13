package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/qazz92/goTest/controller"
	"strconv"
	"crypto/tls"
	"github.com/headzoo/surf"
)

func main() {
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

	r.GET("/getTimeTable", func(c *gin.Context) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		bow := surf.NewBrowser()

		//jar, _ := cookiejar.New(nil)
		//bow.SetCookieJar(jar)
		bow.SetTransport(tr)

		err := bow.Open("https://student.donga.ac.kr/Univ/SUD/SSUD0000.aspx?m=1")

		if err != nil {
			panic(err)
		}

		// Log in to the site.
		login, err := bow.Form("form#frmLogin")
		if err != nil {
			panic(err)
		}
		err = login.Input("txtStudentCd", "1124305")
		if err != nil {
			panic(err)
		}
		err = login.Input("txtPasswd", "Ekfqlc152")
		if err != nil {
			panic(err)
		}




		//err = bow.Open("https://student.donga.ac.kr/Univ/SUD/SSUD0000.aspx?m=1")
		//if err != nil {
		//	panic(err)
		//}


	})

	r.Run(":3000")
}
