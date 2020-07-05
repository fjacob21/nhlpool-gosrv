package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/web"
)

func main() {
	fmt.Println(data.UserHash("fred"))

	log.Println("Creating dummy messages")

	log.Println("Attempting to start HTTP Server.")

	handler := &web.RegexpHandler{}

	handler.HandleFunc("^/player/(.*)/login/$", web.HandlePlayerLoginRequest)
	handler.HandleFunc("^/player/(.*)/logout/$", web.HandlePlayerLogoutRequest)
	handler.HandleFunc("^/player/(.*)/changepassword/$", web.HandlePlayerChangePasswordRequest)
	handler.HandleFunc("^/player/(.*)/$", web.HandlePlayerRequest)
	handler.HandleFunc("^/player/import/$", web.HandlePlayerImportRequest)
	handler.HandleFunc("^/player/$", web.HandlePlayersRequest)
	handler.HandleFunc("^/$", web.HandleRootRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(8080), handler)

	if err != nil {
		log.Printf("Server failed starting. Error: %s", err.Error())
	}
}
