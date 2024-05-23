package controllers

import (
	"brujulavirtual-auth/src/common"
	"brujulavirtual-auth/src/register/domain/models"
	"brujulavirtual-auth/src/register/domain/ports"
	"encoding/json"
	"net/http"
)

type Controller struct {
	service ports.Service
}

func Register(service ports.Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Validate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		controller.ValidatePost(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (controller *Controller) ValidatePost(w http.ResponseWriter, r *http.Request) {

	var auth models.Register
	err := json.NewDecoder(r.Body).Decode(&auth)

	if err != nil {
		common.ErrorResponse(w, "Error processing data", http.StatusBadRequest)
		return
	}

	if !auth.IsValid() {
		common.ErrorResponse(w, "Incomplete or invalid authentication data", http.StatusBadRequest)
		return
	}

	_, err = controller.service.Save(auth)
	if err != nil {
		common.ErrorResponse(w, "Authentication Error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
