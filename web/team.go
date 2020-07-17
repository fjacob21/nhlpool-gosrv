package web

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"nhlpool.com/service/go/nhlpool/controller"
	"nhlpool.com/service/go/nhlpool/data"
)

func getRequestInfo(r *http.Request) (string, string) {
	reg, _ := regexp.Compile(`/league/([^/]*)/team/([^/]*)/.*`)
	results := reg.FindStringSubmatch(r.RequestURI)
	return results[1], results[2]
}

// HandleTeamRequest Handle the web request for league/<league>
func HandleTeamRequest(w http.ResponseWriter, r *http.Request) {
	league, team := getRequestInfo(r)
	log.Println("Team:", r.Method, league, team)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetTeam(league, team))
		break
	case http.MethodDelete:
		var request data.DeleteTeamRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeleteTeam(league, team, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
