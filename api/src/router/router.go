package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRoute() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
