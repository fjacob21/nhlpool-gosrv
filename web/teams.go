package web

import (
	"encoding/json"
	"log"
	"net/http"

	"nhlpool.com/service/go/nhlpool/controller"

	"nhlpool.com/service/go/nhlpool/data"
)

// HandleTeamsRequest Handle the web request for teams
func HandleTeamsRequest(w http.ResponseWriter, r *http.Request) {
	league := getLeague(r)
	log.Println("Team:", r.Method, league)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetTeams(league))
		break
	case http.MethodPost:
		var request data.AddTeamRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			break
		}
		handleSuccess(&w, controller.AddTeam(league, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
