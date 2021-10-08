package handler

import (
	"context"
	"encoding/json"
	"strconv"

	// "fmt"
	// "log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		todo := new(model.CreateTODORequest)
		todo.Subject = r.FormValue("subject")
		todo.Description = r.FormValue("description")
		if todo.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			rsp, err := h.Create(context.TODO(), todo)

			if err != nil {
				// log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
			}

			err = json.NewEncoder(w).Encode(rsp)
			if err != nil {
				// log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	} else if r.Method == http.MethodPut {
		todo := new(model.UpdateTODORequest)
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		todo.ID = id
		todo.Subject = r.FormValue("subject")
		todo.Description = r.FormValue("description")
		if todo.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else if todo.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			rsp, err := h.Update(context.TODO(), todo)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			if rsp.TODO.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
			}
			err = json.NewEncoder(w).Encode(rsp)
			if err != nil {
				// log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}

}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	res, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}
	cres := new(model.CreateTODOResponse)
	cres.TODO = *res
	return cres, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	res, err := h.svc.UpdateTODO(ctx, int64(req.ID), req.Subject, req.Description)
	if err != nil {
		return nil, err
	}
	ures := new(model.UpdateTODOResponse)
	ures.TODO = *res
	return ures, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
