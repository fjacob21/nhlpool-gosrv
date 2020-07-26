package web

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"nhlpool.com/service/go/nhlpool/controller"
	"nhlpool.com/service/go/nhlpool/data"
)

func getRequestMatchupInfo(r *http.Request) (string, int, string) {
	reg, _ := regexp.Compile(`/league/([^/]*)/season/([^/]*)/matchup/([^/]*)/.*`)
	results := reg.FindStringSubmatch(r.RequestURI)
	year, _ := strconv.Atoi(results[2])
	return results[1], year, results[3]
}

// HandleMatchupRequest Handle the web request for league/<league>
func HandleMatchupRequest(w http.ResponseWriter, r *http.Request) {
	league, year, id := getRequestMatchupInfo(r)
	log.Println("Matchup:", r.Method, league, year, id)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetMatchup(league, year, id))
		break
	case http.MethodPost:
		var request data.EditMatchupRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.EditMatchup(league, year, id, request))
		break
	case http.MethodDelete:
		var request data.DeleteMatchupRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeleteMatchup(league, year, id, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
