package main

import (
	"log"
	"wic-server/engine"
	"wic-server/tcpServer"
)

func main() {
	port := "3000"
	games := engine.NewGamesPool()
	if err := tcpServer.Run(port, games); err != nil {
		log.Fatal(err.Error())
	}
}
