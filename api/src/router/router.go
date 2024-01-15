package router

import "github.com/gorilla/mux"

func GenerateRoute() *mux.Router {
	return mux.NewRouter()
}
