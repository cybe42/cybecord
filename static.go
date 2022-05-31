package main

import (
	"embed"
	"mime"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var static embed.FS

type fileHandler struct {
	engine *gin.Engine
}

const (
	defaultFileType = "text/plain; charset=utf-8" // The default file type if we fail to parse from extension.
)

func getFileExtension(filePath string) string {
	var fileSplitted []string = strings.Split(filePath, ".")
	var splittedLength int = len(fileSplitted)
	if splittedLength < 2 {
		return ""
	} else {
		return "." + fileSplitted[splittedLength-1] // get last element from array
	}
}

func getFileType(filePath string) string {
	var fileExtension string = getFileExtension(filePath)
	var fileType string = mime.TypeByExtension(fileExtension)
	if fileType == "" {
		return defaultFileType // return file type if we could not find file extension
	} else {
		return fileType
	}
}

func (handler *fileHandler) addFileHandle(path string, filePath string) {
	file, _ := static.ReadFile(filePath)
	var fileContent string = string(file)

	handler.engine.GET(path, func(c *gin.Context) {
		c.Header("Content-Type", getFileType(filePath))
		c.String(200, fileContent)
	}) // handling static file
}

func Static(r *gin.Engine) {
	var filehandler fileHandler = fileHandler{engine: r}
	staticDir, _ := static.ReadDir("static")
	for _, file := range staticDir { // looping through all the embedded files
		var fileName string = file.Name()
		if fileName == "index.html" {
			filehandler.addFileHandle("/", "static/"+fileName)
		} else if getFileExtension(fileName) == ".html" { // remove .html from path
			filehandler.addFileHandle("/"+fileName[:len(fileName)-len(".html")], "static/"+fileName)
		} else { // else just serve file regularly
			filehandler.addFileHandle("/"+fileName, "static/"+fileName)
		}
	}
}
