package web

import (
	"encoding/json"
	"log"
	"net/http"

	"nhlpool.com/service/go/nhlpool/controller"

	"nhlpool.com/service/go/nhlpool/data"
)

// HandleLeaguesRequest Handle the web request for player
func HandleLeaguesRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Leagues:", r.Method)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetLeagues())
		break
	case http.MethodPost:
		var request data.AddLeagueRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			break
		}
		handleSuccess(&w, controller.AddLeague(request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
