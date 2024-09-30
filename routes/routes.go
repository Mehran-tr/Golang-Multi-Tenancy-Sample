package routes

import (
	"multi-tenant-go-app/controllers"
	"multi-tenant-go-app/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Apply tenant middleware to user routes
	r.Handle("/users", middleware.TenantMiddleware(http.HandlerFunc(controllers.GetUsers)))

	return r
}
