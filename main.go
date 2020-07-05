package main

import (
	"log"
	"net/http"
	"strconv"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/web"
)

func main() {
	configs := LoadConfigs()
	admin := store.GetStore().GetPlayer(configs.Admin.ID)
	if admin != nil {
		store.GetStore().DeletePlayer(admin)
	}
	store.GetStore().AddPlayer(&configs.Admin)
	log.Println("Attempting to start HTTP Server.")

	handler := &web.RegexpHandler{}

	handler.HandleFunc("^/player/(.*)/login/$", web.HandlePlayerLoginRequest)
	handler.HandleFunc("^/player/(.*)/logout/$", web.HandlePlayerLogoutRequest)
	handler.HandleFunc("^/player/(.*)/changepassword/$", web.HandlePlayerChangePasswordRequest)
	handler.HandleFunc("^/player/(.*)/$", web.HandlePlayerRequest)
	handler.HandleFunc("^/player/import/$", web.HandlePlayerImportRequest)
	handler.HandleFunc("^/player/$", web.HandlePlayersRequest)
	handler.HandleFunc("^/$", web.HandleRootRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(configs.Port), handler)

	if err != nil {
		log.Printf("Server failed starting. Error: %s", err.Error())
	}
}
