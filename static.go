package main

import (
	"embed"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var static embed.FS

/*var (
	Index     string
	Bootstrap string
)*/ // No need to be global for now

func Static(r *gin.Engine) {

	index, _ := static.ReadFile("static/index.html")
	bootstrap, _ := static.ReadFile("static/bootstrap.min.css")

	var Index string = string(index)
	var Bootstrap string = string(bootstrap)

	// Serve the static files.
	r.GET("/", func(c *gin.Context) { c.String(200, Index) })
	r.GET("/bootstrap.min.js", func(c *gin.Context) { c.String(200, Bootstrap) })
}
