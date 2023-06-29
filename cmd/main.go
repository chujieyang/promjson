package main

import (
	"fmt"
	"net/http"

	"github.com/chujieyang/promjson/prom"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/metricsjson", func(c *gin.Context) {
		prom_metrics_url := c.DefaultQuery("url", "")
		jsontext, err := prom.Promjson(prom_metrics_url)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("convert to json exception: %+v", err))
			return
		}
		c.String(http.StatusOK, jsontext)
	})
	r.Run(":19101")
}
