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

type fileHandler struct {
	engine *gin.Engine
}

func (handler *fileHandler) addFileHandle(path string, filePath string) {
	file, _ := static.ReadFile(filePath)
	var fileContent string = string(file)
	handler.engine.GET(path, func(c *gin.Context) { c.String(200, fileContent) })
}

func Static(r *gin.Engine) {
	var filehandler fileHandler = fileHandler{engine: r}

	filehandler.addFileHandle("/", "static/index.html")
	filehandler.addFileHandle("/bootstrap.min.css", "static/bootstrap.min.css")
}
