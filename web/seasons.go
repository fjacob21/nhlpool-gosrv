package web

import (
	"encoding/json"
	"net/http"

	"nhlpool.com/service/go/nhlpool/controller"

	"nhlpool.com/service/go/nhlpool/data"
)

// HandleSeasonsRequest Handle the web request for teams
func HandleSeasonsRequest(w http.ResponseWriter, r *http.Request) {
	league := getLeague(r)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetSeasons(league))
		break
	case http.MethodPost:
		var request data.AddSeasonRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			break
		}
		handleSuccess(&w, controller.AddSeason(league, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
