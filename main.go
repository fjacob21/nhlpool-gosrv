package main

import (
	"log"
	"net/http"
	"strconv"

	"nhlpool.com/service/go/nhlpool/web"
)

func main() {
	log.Println("Creating dummy messages")

	log.Println("Attempting to start HTTP Server.")

	handler := &web.RegexpHandler{}

	handler.HandleFunc("/", web.HandleRootRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(8080), handler)

	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}
}
