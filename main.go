package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Static(r) // Handle everything static

	r.Run()
}
