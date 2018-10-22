package routes

import (
	"go-basic-rest-api/handlers"
	"go-basic-rest-api/utils"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	AuthFirst   bool
	AuthCookie  bool
	RouteHandle handlers.HandlerRoute
}

type Routes []Route

var DefinedRoutes = Routes{
	Route{
		"GetAllTodos",
		"GET",
		"/api/todos",
		false,
		false,
		handlers.GetAllTodos,
	},
	Route{
		"CreateTodo",
		"POST",
		"/api/todos",
		true,
		false,
		handlers.CreateTodo,
	},
	Route{
		"GetTodo",
		"GET",
		"/api/todos/{id}",
		false,
		false,
		handlers.GetTodo,
	},
	Route{
		"UpdateTodo",
		"PUT",
		"/api/todos/{id}",
		true,
		false,
		handlers.UpdateTodo,
	},
	Route{
		"DeleteTodo",
		"Delete",
		"/api/todos/{id}",
		true,
		false,
		handlers.DeleteTodo,
	},
	Route{
		"CreateUser",
		"POST",
		"/api/users",
		false,
		false,
		handlers.CreateUser,
	},
	Route{
		"GetAllUsers",
		"GET",
		"/api/users",
		false,
		false,
		handlers.GetAllUsers,
	},
	Route{
		"LoginUsers",
		"POST",
		"/api/login",
		false,
		false,
		handlers.AuthUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/api/users/{id}",
		false,
		false,
		handlers.GetUser,
	},
	Route{
		"Upload User Photo",
		"POST",
		"/api/user-photos",
		true,
		false,
		handlers.UploadUserPhoto,
	},
	Route{
		"Get User Photo",
		"GET",
		"/api/user-photos",
		true,
		false,
		handlers.GetUserPhoto,
	},
	Route{
		"Upload Photo Profile",
		"POST",
		"/api/upload-photo-profiles",
		true,
		false,
		handlers.UploadPhotoProfile,
	},
	Route{
		"RenderIndex",
		"GET",
		"/",
		false,
		true,
		handlers.RenderIndex,
	},
	Route{
		"RenderHome",
		"GET",
		"/home",
		false,
		true,
		handlers.RenderHome,
	},
}

func NotFoundRoute(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.RequestURI, "/api/") {
		handlers.ProcessJSON(w, http.StatusNotFound, []byte(`"NOT FOUND"`), utils.DATA_NOT_FOUND)
	} else {
		handlers.RenderNotFound(w, r)
	}
}
