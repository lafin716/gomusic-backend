package main

import (
	"github.com/lafin716/gomusic/rest"
	"log"
)

func main() {
	log.Println("Main Log...")
	log.Fatal(rest.RunApi("127.0.0.1:8000"))
}