package web

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler Object that handle regex pattern
type RegexpHandler struct {
	routes []*route
}

// Handler Add an handler
func (h *RegexpHandler) Handler(pattern string, handler http.Handler) {
	reg, _ := regexp.Compile(pattern)
	h.routes = append(h.routes, &route{reg, handler})
}

// HandleFunc Add an handler function
func (h *RegexpHandler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	reg, _ := regexp.Compile(pattern)
	h.routes = append(h.routes, &route{reg, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}
