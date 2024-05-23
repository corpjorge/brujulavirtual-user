package routes

import (
	"brujulavirtual-auth/src/register/infrastructure/controllers"
	"net/http"
)

func Router(register controllers.Controller, mux *http.ServeMux) {
	mux.HandleFunc("/register", register.Validate)
}
