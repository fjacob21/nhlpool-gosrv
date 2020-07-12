package web

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"nhlpool.com/service/go/nhlpool/controller"
	"nhlpool.com/service/go/nhlpool/data"
)

func getLeague(r *http.Request) string {
	reg, _ := regexp.Compile(`/league/([^/]*)/.*`)
	league := reg.FindStringSubmatch(r.RequestURI)[1]
	return league
}

// HandleLeagueRequest Handle the web request for league/<league>
func HandleLeagueRequest(w http.ResponseWriter, r *http.Request) {
	league := getLeague(r)
	log.Println("League:", r.Method, league)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetLeague(league))
		break
	case http.MethodDelete:
		var request data.DeleteLeagueRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeleteLeague(league, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
