package web

import (
	"encoding/json"
	"log"
	"net/http"

	"nhlpool.com/service/go/nhlpool/controller"

	"nhlpool.com/service/go/nhlpool/data"
)

// HandleWinnersRequest Handle the web request for winners
func HandleWinnersRequest(w http.ResponseWriter, r *http.Request) {
	league, year := getSeasonInfo(r)
	log.Println("Winner:", r.Method, league, year)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetWinners(league, year))
		break
	case http.MethodPost:
		var request data.AddWinnerRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			break
		}
		handleSuccess(&w, controller.AddWinner(league, year, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
