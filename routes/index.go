package routes

import (
	"go-basic-rest-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	RouteHandle handlers.HandlerRoute
}

type Routes []Route

var DefinedRoutes = Routes{
	Route{
		"GetAllTodos",
		"GET",
		"/todos",
		handlers.GetAllTodos,
	},
	Route{
		"CreateTodo",
		"POST",
		"/todos",
		handlers.CreateTodo,
	},
	Route{
		"GetTodo",
		"GET",
		"/todos/{id}",
		handlers.GetTodo,
	},
	Route{
		"UpdateTodo",
		"PUT",
		"/todos/{id}",
		handlers.UpdateTodo,
	},
	Route{
		"DeleteTodo",
		"Delete",
		"/todos/{id}",
		handlers.DeleteTodo,
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		handlers.CreateUser,
	},
	Route{
		"GetAllUsers",
		"GET",
		"/users",
		handlers.GetAllUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/users/{id}",
		handlers.GetUser,
	},
}
