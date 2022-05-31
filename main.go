package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Static()

	r.GET("/", func(c *gin.Context) { c.String(200, Index) })

	r.Run()
}
