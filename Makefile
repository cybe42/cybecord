source = main.go static.go flags.go

build:
	go build -o bin/cybecord-debug $(source) 

release:
	go build -o bin/cybecord -ldflags "-s -w" $(source) release.go

run:
	go run $(source)