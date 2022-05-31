package main

import "embed"

//go:embed static/*
var static embed.FS

var (
	Index     string
	Bootstrap string
)

func Static() {
	index, _ := static.ReadFile("static/index.html")
	bootstrap, _ := static.ReadFile("static/bootstrap.min.css")

	Index = string(index)
	Bootstrap = string(bootstrap)
}
