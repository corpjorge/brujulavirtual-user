package register

import (
	"brujulavirtual-auth/src/register/application/services"
	"brujulavirtual-auth/src/register/infrastructure/controllers"
	"brujulavirtual-auth/src/register/infrastructure/repositories"
	"brujulavirtual-auth/src/register/infrastructure/routes"
	"net/http"
)

func Module(mux *http.ServeMux) {
	registerRepository := repositories.Register()
	registerService := services.Register(registerRepository)
	registerController := controllers.Register(registerService)

	routes.Router(*registerController, mux)
}
