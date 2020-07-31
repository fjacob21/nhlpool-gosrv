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

func getRequestStandingInfo(r *http.Request) (string, int, string) {
	reg, _ := regexp.Compile(`/league/([^/]*)/season/([^/]*)/standing/([^/]*)/.*`)
	results := reg.FindStringSubmatch(r.RequestURI)
	year, _ := strconv.Atoi(results[2])
	return results[1], year, results[3]
}

// HandleStandingRequest Handle the web request for league/<league>
func HandleStandingRequest(w http.ResponseWriter, r *http.Request) {
	league, year, team := getRequestStandingInfo(r)
	log.Println("Standing:", r.Method, league, year, team)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetStanding(league, year, team))
		break
	case http.MethodPost:
		var request data.EditStandingRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.EditStanding(league, year, team, request))
		break
	case http.MethodDelete:
		var request data.DeleteStandingRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeleteStanding(league, year, team, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
