package web

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"nhlpool.com/service/go/nhlpool/controller"
	"nhlpool.com/service/go/nhlpool/data"
)

func getPlayer(r *http.Request) string {
	reg, _ := regexp.Compile(`/player/([^/]*)/.*`)
	player := reg.FindStringSubmatch(r.RequestURI)[1]
	return player
}

// HandlePlayerRequest Handle the web request for player/<player>
func HandlePlayerRequest(w http.ResponseWriter, r *http.Request) {
	player := getPlayer(r)
	log.Println("Player:", r.Method, player)
	switch r.Method {
	case http.MethodGet:
		handleSuccess(&w, controller.GetPlayer(player))
		break
	case http.MethodPost:
		var request data.EditPlayerRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.EditPlayer(player, request))
		break
	case http.MethodDelete:
		var request data.DeletePlayerRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.DeletePlayer(player, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

// HandlePlayerLoginRequest Handle the web request for player/<player>/login
func HandlePlayerLoginRequest(w http.ResponseWriter, r *http.Request) {
	player := getPlayer(r)
	log.Println("Player login:", r.Method, player)
	switch r.Method {
	case http.MethodPost:
		var request data.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.Login(player, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

// HandlePlayerLogoutRequest Handle the web request for player/<player>/logout
func HandlePlayerLogoutRequest(w http.ResponseWriter, r *http.Request) {
	player := getPlayer(r)
	log.Println("Player logout:", r.Method, player)
	switch r.Method {
	case http.MethodPost:
		var request data.LogoutRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.Logout(player, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

// HandlePlayerChangePasswordRequest Handle the web request for player/<player>/changepassword
func HandlePlayerChangePasswordRequest(w http.ResponseWriter, r *http.Request) {
	player := getPlayer(r)
	log.Println("Player change password:", r.Method, player)
	switch r.Method {
	case http.MethodPost:
		var request data.ChangePasswordRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handleSuccess(&w, controller.ChangePassword(player, request))
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
