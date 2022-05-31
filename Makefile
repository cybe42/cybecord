source = main.go static.go

build:
	go build -o bin/cybecord-debug $(source) 

release:
	go build -o bin/cybecord -ldflags "-s -w" $(source) release.go

