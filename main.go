package main

//
import (
	"./logs"
	"./web"
	"flag"
)

var (
	port    string
	pattern string //api  cmd
)

func initFlag() {
	flag.StringVar(&port, "port", "8002", "server port")
	flag.StringVar(&pattern, "p", "api", "runing pattern")
	flag.Parse()
}

func main() {
	initFlag()
	//
	logs.Log.Informational("pattern: %s", pattern)
	if pattern == "api" {
		web.Run(port)
	}
}
