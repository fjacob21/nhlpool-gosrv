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

func getSeasonInfo(r *http.Request) (string, int) {
	reg, _ := regexp.Compile(`/league/([^/]*)/season/([^/]*)/.*`)
	results := reg.FindStringSubmatch(r.RequestURI)
	year, _ := strconv.Atoi(results[2])
	return results[1], year
}

// HandleSeasonRequest Handle the web request for league/<league>
func HandleSeasonRequest(w http.ResponseWriter, r *http.Request) {
	league, year := getSeasonInfo(r)
	log.Println("Season:", r.Method, league, year)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetSeason(league, year))
		break
	case http.MethodDelete:
		var request data.DeleteSeasonRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeleteSeason(league, year, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
