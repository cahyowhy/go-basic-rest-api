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
		"/api/todos",
		handlers.GetAllTodos,
	},
	Route{
		"CreateTodo",
		"POST",
		"/api/todos",
		handlers.CreateTodo,
	},
	Route{
		"GetTodo",
		"GET",
		"/api/todos/{id}",
		handlers.GetTodo,
	},
	Route{
		"UpdateTodo",
		"PUT",
		"/api/todos/{id}",
		handlers.UpdateTodo,
	},
	Route{
		"DeleteTodo",
		"Delete",
		"/api/todos/{id}",
		handlers.DeleteTodo,
	},
	Route{
		"CreateUser",
		"POST",
		"/api/users",
		handlers.CreateUser,
	},
	Route{
		"GetAllUsers",
		"GET",
		"/api/users",
		handlers.GetAllUsers,
	},
	Route{
		"LoginUsers",
		"POST",
		"/api/login",
		handlers.AuthUsers,
	},
	Route{
		"Authorization",
		"GET",
		"/api/header-check",
		handlers.CekHeaderAuth,
	},
	Route{
		"GetUser",
		"GET",
		"/api/users/{id}",
		handlers.GetUser,
	},
	Route{
		"RenderIndex",
		"GET",
		"/",
		handlers.RenderIndex,
	},
}
