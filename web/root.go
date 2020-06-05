package web

import (
	"log"
	"net/http"

	"nhlpool.com/service/go/nhlpool/data"
)

// HandleRootRequest Handle the web request for root
func HandleRootRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r.Method)
	switch r.Method {
	case http.MethodGet:
		res := data.RootReply{Name: "nhlpool", Version: "v1.0"}
		handleSuccess(&w, res)
		break
	default:
		handleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
