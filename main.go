package main

//
import (
	//"./logs"
	"./web"
	"flag"
)

var (
	port string
)

func initFlag() {
	flag.StringVar(&port, "port", "8002", "server port")
	flag.Parse()
}

func main() {
	initFlag()
	web.Run(port)
}
