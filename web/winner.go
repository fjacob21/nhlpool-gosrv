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

func getRequestWinnerInfo(r *http.Request) (string, int, string) {
	reg, _ := regexp.Compile(`/league/([^/]*)/season/([^/]*)/winner/([^/]*)/.*`)
	results := reg.FindStringSubmatch(r.RequestURI)
	year, _ := strconv.Atoi(results[2])
	return results[1], year, results[3]
}

// HandleWinnerRequest Handle the web request for league/<league>
func HandleWinnerRequest(w http.ResponseWriter, r *http.Request) {
	league, year, player := getRequestWinnerInfo(r)
	log.Println("Winner:", r.Method, league, year, player)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetWinner(league, year, player))
		break
	case http.MethodPost:
		var request data.EditWinnerRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.EditWinner(league, year, player, request))
		break
	case http.MethodDelete:
		var request data.DeleteWinnerRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeleteWinner(league, year, player, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
