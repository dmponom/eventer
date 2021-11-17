package httpserver

import "net/http"

type Server interface {
	Register()
}

type API interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}
