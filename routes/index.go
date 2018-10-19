package routes

import (
	"go-basic-rest-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	AuthFirst   bool
	RouteHandle handlers.HandlerRoute
}

type Routes []Route

var DefinedRoutes = Routes{
	Route{
		"GetAllTodos",
		"GET",
		"/api/todos",
		false,
		handlers.GetAllTodos,
	},
	Route{
		"CreateTodo",
		"POST",
		"/api/todos",
		true,
		handlers.CreateTodo,
	},
	Route{
		"GetTodo",
		"GET",
		"/api/todos/{id}",
		false,
		handlers.GetTodo,
	},
	Route{
		"UpdateTodo",
		"PUT",
		"/api/todos/{id}",
		true,
		handlers.UpdateTodo,
	},
	Route{
		"DeleteTodo",
		"Delete",
		"/api/todos/{id}",
		true,
		handlers.DeleteTodo,
	},
	Route{
		"CreateUser",
		"POST",
		"/api/users",
		false,
		handlers.CreateUser,
	},
	Route{
		"GetAllUsers",
		"GET",
		"/api/users",
		false,
		handlers.GetAllUsers,
	},
	Route{
		"LoginUsers",
		"POST",
		"/api/login",
		false,
		handlers.AuthUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/api/users/{id}",
		false,
		handlers.GetUser,
	},
	Route{
		"UploadTodo",
		"POST",
		"/api/upload-todos",
		true,
		handlers.UploadTodo,
	},
	Route{
		"RenderIndex",
		"GET",
		"/",
		false,
		handlers.RenderIndex,
	},
}