package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/qazz92/GoBoo/controller"
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

	r.Run(":3000")
}
