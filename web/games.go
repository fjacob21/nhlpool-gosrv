package web

import (
	"encoding/json"
	"log"
	"net/http"

	"nhlpool.com/service/go/nhlpool/controller"

	"nhlpool.com/service/go/nhlpool/data"
)

// HandleGamesRequest Handle the web request for game
func HandleGamesRequest(w http.ResponseWriter, r *http.Request) {
	league, year := getSeasonInfo(r)
	log.Println("Game:", r.Method, league, year)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetGames(league, year))
		break
	case http.MethodPost:
		var request data.AddGameRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			break
		}
		handleSuccess(&w, controller.AddGame(league, year, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

// HandleGameUpdateRequest Handle the web request for game/update
func HandleGameUpdateRequest(w http.ResponseWriter, r *http.Request) {
	league, year := getSeasonInfo(r)
	log.Println("Game update:", r.Method, league, year)
	switch r.Method {
	case http.MethodPost:
		var request data.EditGameRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.UpdateGame(league, year, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
