package httphandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *methods) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := &HealthCheckResponse{Message: "ok"}

	respMsg, err := json.Marshal(resp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respMsg)
}
