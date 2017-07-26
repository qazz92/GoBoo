package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/qazz92/GoBoo/controller"
	"runtime"
	"github.com/jpillora/ipfilter"
	"github.com/qazz92/GoBoo/realip"
	"log"
	"github.com/gin-contrib/gzip"
)
func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func BlockMiddleware(c *gin.Context) {
	host := realip.RealIP(c.Request)
	f, _ := ipfilter.New(ipfilter.Options{
		BlockByDefault: true,
	})
	f.AllowCountry("KR")
	log.Println(host)
	if f.Blocked(host) {
		c.Abort()
		return
	}
	// Pass on to the next-in-chain
	c.Next()
}

func main() {
	runtime.GOMAXPROCS(MaxParallelism())
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))


	r.Use(BlockMiddleware)
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
