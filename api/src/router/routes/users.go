package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Func:                  controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Func:                  controllers.SearchUsers,
		RequireAuthentication: false,
	}, {
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Func:                  controllers.UpdateUser,
		RequireAuthentication: false,
	}, {
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Func:                  controllers.DeleteUser,
		RequireAuthentication: false,
	}, {
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Func:                  controllers.SearchUser,
		RequireAuthentication: false,
	},
}
