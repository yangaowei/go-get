package main

//
import (
	"./logs"
	"./web"
)

func main() {
	logs.Log.Informational("statrt server")
	web.Run()
}
