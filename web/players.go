package web

import (
	"encoding/json"
	"log"
	"net/http"

	"nhlpool.com/service/go/nhlpool/controller"

	"nhlpool.com/service/go/nhlpool/data"
)

// HandlePlayersRequest Handle the web request for player
func HandlePlayersRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Players:", r.Method)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetPlayers())
		break
	case http.MethodPost:
		var request data.AddPlayerRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			break
		}
		handleSuccess(&w, controller.AddPlayer(request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

// HandlePlayerImportRequest Handle the web request for player/import
func HandlePlayerImportRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Players import:", r.Method)
	switch r.Method {
	case http.MethodPost:
		var request data.ImportPlayerRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.ImportPlayer(request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
