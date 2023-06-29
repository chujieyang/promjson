package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chujieyang/promjson/prom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	help bool
	port int
)

func init() {
	flag.BoolVar(&help, "h", false, "help infomation")
	flag.IntVar(&port, "p", 19100, "set port of server running")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stdout, `promjson: version 1.0.0
Usage: promjson [-h] [-p port]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	if help {
		flag.Usage()
		return
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/metricsjson", func(c *gin.Context) {
		prom_metrics_url := c.DefaultQuery("url", "")
		jsontext, err := prom.Promjson(prom_metrics_url)
		if err != nil {
			c.JSON(200, gin.H{"code": -1, "message": err.Error(), "data": nil})
			return
		}
		c.JSON(200, gin.H{"code": 0, "message": "success", "data": jsontext})
	})

	r.Run(fmt.Sprintf(":%d", port))
}
