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

func getRequestPredictionInfo(r *http.Request) (string, int, string, string) {
	reg, _ := regexp.Compile(`/league/([^/]*)/season/([^/]*)/prediction/([^/]*)/([^/]*)/.*`)
	results := reg.FindStringSubmatch(r.RequestURI)
	year, _ := strconv.Atoi(results[2])
	return results[1], year, results[3], results[4]
}

// HandlePredictionRequest Handle the web request for league/<league>
func HandlePredictionRequest(w http.ResponseWriter, r *http.Request) {
	league, year, player, matchup := getRequestPredictionInfo(r)
	log.Println("Prediction:", r.Method, league, year, player, matchup)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetPrediction(league, year, player, matchup))
		break
	case http.MethodPost:
		var request data.EditPredictionRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.EditPrediction(league, year, player, matchup, request))
		break
	case http.MethodDelete:
		var request data.DeletePredictionRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeletePrediction(league, year, player, matchup, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
