package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miles990/ratelimiter-example/global"
	"github.com/miles990/ratelimiter-example/ratelimiter"
)

// var (
// 	limiter map[string]*ratelimiter
// )

var (
	help = flag.Bool("help", false, "Show this help")
	port = flag.Int("port", 80, "http server listen port")
	t    = flag.Int("t", 60, "limit time (second)")
	num  = flag.Int("num", 60, "limit num")
	// configFile = flag.String("conf", "config.yaml", "The server configurate file")
)

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	limiter := ratelimiter.NewRateLimiter(
		ratelimiter.LimitTime(time.Duration(*t)*time.Second),
		ratelimiter.LimitNum(*num),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		_, _, err := limiter.Check(c.ClientIP())

		if err != nil {
			c.JSON(403, gin.H{
				"count": global.Num(),
				"err":   err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"count": global.Num(),
			"info":  global.GetAllIPInfos(),
			"time":  time.Now(),
		})
	})
	router.Run(fmt.Sprintf(":%d", *port))
}
