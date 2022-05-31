package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	flags := ParseFlags()

	Static(r) // Handle everything static

	r.Run(":" + strconv.Itoa(flags.port))
}
