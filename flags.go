package main

import "flag"

type Flags struct {
	port int
}

func ParseFlags() Flags {
	port := flag.Int("port", 8080, "Port used for listening. Default is 8080")

	flag.Parse()

	return Flags{
		port: *port,
	}
}
