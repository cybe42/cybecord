package main

import (
	"embed"
	"mime"
	"strings"

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

const (
	defaultFileType = "text/plain; charset=utf-8"
)

func getFileType(filePath string) string {
	var fileSplitted []string = strings.Split(filePath, ".")
	var splittedLength int = len(fileSplitted)
	if splittedLength < 2 {
		return defaultFileType
	} else {
		var fileType string = mime.TypeByExtension("." + fileSplitted[splittedLength-1]) // get last element from array
		if fileType == "" {
			return defaultFileType
		} else {
			return fileType
		}
	}
}

func (handler *fileHandler) addFileHandle(path string, filePath string) {
	file, _ := static.ReadFile(filePath)
	var fileContent string = string(file)

	handler.engine.GET(path, func(c *gin.Context) {
		c.Header("Content-Type", getFileType(filePath))
		c.String(200, fileContent)
	})
}

func Static(r *gin.Engine) {
	var filehandler fileHandler = fileHandler{engine: r}
	staticDir, _ := static.ReadDir("static")
	for _, file := range staticDir {
		var fileName string = file.Name()
		if fileName != "index.html" {
			filehandler.addFileHandle("/"+fileName, "static/"+fileName)
		} else {
			filehandler.addFileHandle("/", "static/"+fileName)
		}
	}
}
