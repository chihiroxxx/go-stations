package handler

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rsp := &model.HealthzResponse{Message: "OK"}
	err := json.NewEncoder(w).Encode(rsp)
	if err != nil {
		log.Println(err)
	}
	// j, _ := json.Marshal(rsp)
	// fmt.Fprintf(w, string(j))
}
