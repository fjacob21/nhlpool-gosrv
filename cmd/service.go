package cmd

import (
	"log"
	"net/http"
	"strconv"

	"nhlpool.com/service/go/nhlpool/config"
	"nhlpool.com/service/go/nhlpool/store"
	"nhlpool.com/service/go/nhlpool/web"
)

// Service Execute the nhlpoool service
func Service() {
	configs := config.LoadConfigs()
	admin := store.GetStore().Player().GetPlayer(configs.Admin.ID)
	if admin != nil {
		store.GetStore().Player().DeletePlayer(admin)
	}
	store.GetStore().Player().AddPlayer(&configs.Admin)
	log.Println("Attempting to start HTTP Server.")

	handler := &web.RegexpHandler{}

	handler.HandleFunc("^/league/(.*)/season/(.*)/standing/$", web.HandleStandingsRequest)
	handler.HandleFunc("^/league/(.*)/season/(.*)/standing/(.*)/$", web.HandleStandingRequest)
	handler.HandleFunc("^/league/(.*)/season/$", web.HandleSeasonsRequest)
	handler.HandleFunc("^/league/(.*)/season/(.*)/$", web.HandleSeasonRequest)
	handler.HandleFunc("^/league/(.*)/team/$", web.HandleTeamsRequest)
	handler.HandleFunc("^/league/(.*)/team/(.*)/$", web.HandleTeamRequest)
	handler.HandleFunc("^/league/$", web.HandleLeaguesRequest)
	handler.HandleFunc("^/league/(.*)/$", web.HandleLeagueRequest)
	handler.HandleFunc("^/player/import/$", web.HandlePlayerImportRequest)
	handler.HandleFunc("^/player/$", web.HandlePlayersRequest)
	handler.HandleFunc("^/player/(.*)/login/$", web.HandlePlayerLoginRequest)
	handler.HandleFunc("^/player/(.*)/logout/$", web.HandlePlayerLogoutRequest)
	handler.HandleFunc("^/player/(.*)/changepassword/$", web.HandlePlayerChangePasswordRequest)
	handler.HandleFunc("^/player/(.*)/$", web.HandlePlayerRequest)
	handler.HandleFunc("^/$", web.HandleRootRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(configs.Port), handler)

	if err != nil {
		log.Printf("Server failed starting. Error: %s", err.Error())
	}
}
