package main

import (
	"log"
	"univer/avito-parser/rest"
)

func main() {
	err := rest.StartServer(8080)
	if err != nil {
		log.Fatal(err)
	}
}
